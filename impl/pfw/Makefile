SHELL := $(SHELL) -o 'pipefail'

all: build
.PHONY: all

build: sapsigner.out
.PHONY: build

clean:
	rm -Rf *.out .build .swiftpm Package.resolved
.PHONY: clean

sapsigner.out: AppleMediaServices.xcframework/* AppleMediaServices.xcframework/macos-arm64_x86_64/* AppleMediaServices.xcframework/macos-arm64_x86_64/AppleMediaServices.framework/* AppleMediaServices.xcframework/macos-arm64_x86_64/AppleMediaServices.framework/Headers/* AppleMediaServices.xcframework/macos-arm64_x86_64/AppleMediaServices.framework/Modules/* CommerceKit.xcframework/* CommerceKit.xcframework/macos-arm64_x86_64/* CommerceKit.xcframework/macos-arm64_x86_64/CommerceKit.framework/* CommerceKit.xcframework/macos-arm64_x86_64/CommerceKit.framework/Headers/* CommerceKit.xcframework/macos-arm64_x86_64/CommerceKit.framework/Modules/* Package.swift Sources/SAPSignerBin/* Sources/SAPSignerLib/* Sources/SAPSignerLib/include/*
	swift build -c release --arch arm64 --arch x86_64
	ln -Lf .build/apple/Products/Release/sapsigner $@
