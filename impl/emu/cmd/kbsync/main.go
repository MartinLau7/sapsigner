package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/emulator"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/guid"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := Main(ctx, flagOutput, flagDSID); err != nil {
		panic(err)
	}
}

func Main(ctx context.Context, oName string, dsid uint64) error {
	oFile, err := openOr(oName, os.Stdout)
	if err != nil {
		return err
	}
	defer oFile.Close()

	artifacts, err := library.Fetch(ctx)
	if err != nil {
		return err
	}

	corefpicxsO, err := library.NewCoreFPICXSObject(artifacts["CoreFP.icxs"])
	if err != nil {
		return err
	}

	corefpO, err := library.NewCoreFPObject(artifacts["CoreFP"])
	if err != nil {
		return err
	}

	commercecoreO, err := library.NewCommerceCoreObject(artifacts["CommerceCore"])
	if err != nil {
		return err
	}

	commercekitO, err := library.NewCommerceKitObject(artifacts["CommerceKit"])
	if err != nil {
		return err
	}

	storeagentO, err := library.NewStoreAgentObject(artifacts["storeagent"])
	if err != nil {
		return err
	}

	e, err := emulator.NewEmulator(corefpicxsO, corefpO, commercecoreO, commercekitO, storeagentO)
	if err != nil {
		return err
	}

	id, err := guid.Get()
	if err != nil {
		return err
	}

	var hwInfo definitions.FairPlayHwInfo
	hwInfo.SetId(id)

	globalContext, err := e.FairPlayGlobalContextInit(&hwInfo)
	if err != nil {
		return err
	}

	oBytes, err := e.FairPlayKBSyncDataWithDSID(globalContext, dsid)
	if err != nil {
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
