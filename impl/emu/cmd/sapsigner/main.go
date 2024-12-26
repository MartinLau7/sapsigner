package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/certificate"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/emulator"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/guid"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/play"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := Main(ctx, flagInput, flagOutput, flagPrimed); err != nil {
		panic(err)
	}
}

func Main(ctx context.Context, iName string, oName string, primeData bool) error {
	iFile, err := openOr(iName, os.Stdin)
	if err != nil {
		return err
	}
	defer iFile.Close()

	oFile, err := openOr(oName, os.Stdout)
	if err != nil {
		return err
	}
	defer iFile.Close()

	id, err := guid.Get()
	if err != nil {
		return err
	}

	crt, err := certificate.Fetch(ctx)
	if err != nil {
		return err
	}

	lib, err := library.Fetch(ctx)
	if err != nil {
		return err
	}

	o, err := library.NewObject(lib)
	if err != nil {
		return err
	}

	e, err := emulator.NewEmulator(o)
	if err != nil {
		return err
	}

	var hwInfo definitions.FairPlayHwInfo
	hwInfo.SetId(id)

	ctxRef, err := e.FairPlaySAPInit(&hwInfo)
	if err != nil {
		return err
	}

	xVer := definitions.FairPlaySAPExchangeVersionRegular
	if primeData {
		xVer = definitions.FairPlaySAPExchangeVersionPrime
	}

	oBuf, returnCode0, err := e.FairPlaySAPExchange(xVer, &hwInfo, ctxRef, crt)
	if err != nil {
		return err
	}
	if returnCode0 != 1 {
		return fmt.Errorf("FairPlaySAPExchange: %d != 1", returnCode0)
	}

	iBuf, err := play.SignSAPSetup(ctx, oBuf)
	if err != nil {
		return err
	}

	_, returnCode1, err := e.FairPlaySAPExchange(xVer, &hwInfo, ctxRef, iBuf)
	if err != nil {
		return err
	}
	if returnCode1 != 0 {
		return fmt.Errorf("FairPlaySAPExchange: %d != 0", returnCode1)
	}

	iBytes, err := io.ReadAll(iFile)
	if err != nil {
		return err
	}

	processData := SignData
	if primeData {
		processData = PrimeData
	}

	oBytes, err := processData(e, ctxRef, iBytes)
	if err != nil {
		return err
	}

	if err := e.FairPlaySAPTeardown(ctxRef); err != nil {
		return err
	}

	if err := e.Close(); err != nil {
		return err
	}

	if _, err := oFile.Write(oBytes); err != nil {
		return err
	}

	return nil
}

func SignData(e *emulator.Emulator, ctxRef *definitions.FPSAPContextOpaqueRef, iBytes []byte) ([]byte, error) {
	return e.FairPlaySAPSign(ctxRef, iBytes)
}

func PrimeData(e *emulator.Emulator, ctxRef *definitions.FPSAPContextOpaqueRef, iBytes []byte) ([]byte, error) {
	return e.FairPlaySAPPrime(ctxRef, definitions.FairPlaySAPPrimeVersionRegular, iBytes)
}
