package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/log"
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

	if err := Main(ctx); err != nil {
		panic(err)
	}
}

func Main(ctx context.Context) error {
	iFile := os.Stdin
	oFile := os.Stdout

	for i, arg := range os.Args {
		var fd **os.File
		switch i {
		case 0:
			continue
		case 1:
			fd = &iFile
		case 2:
			fd = &oFile
		default:
			return errors.New("too many arguments provided")
		}

		if arg == "-" {
			continue
		}

		f, err := os.Open(arg)
		if err != nil {
			return err
		}
		defer f.Close()

		*fd = f
	}

	log.Logger().SetOutput(io.Discard)

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

	oBuf, returnCode0, err := e.FairPlaySAPExchange(definitions.FairPlaySAPExchangeVersionRegular, &hwInfo, ctxRef, crt)
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

	_, returnCode1, err := e.FairPlaySAPExchange(definitions.FairPlaySAPExchangeVersionRegular, &hwInfo, ctxRef, iBuf)
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

	oBytes, err := e.FairPlaySAPSign(ctxRef, iBytes)
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
