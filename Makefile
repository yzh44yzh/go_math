.PHONY: build fmt vet test run android android-install

# The first release targets X11 through XWayland on Linux.
GO_TAGS ?= x11

# Android packaging uses the fyne CLI and a local Android SDK.
ANDROID_HOME ?= $(HOME)/dev/android-sdk
ADB ?= $(ANDROID_HOME)/platform-tools/adb
FYNE ?= $(shell go env GOPATH)/bin/fyne
APK ?= Go_Math.apk

build: fmt vet
	go build -tags $(GO_TAGS) -o go_math .

fmt:
	go fmt ./...

vet:
	go vet -tags $(GO_TAGS) ./...

test:
	go test -tags ci ./...

run:
	go run -tags $(GO_TAGS) .

android: fmt vet
	ANDROID_HOME=$(ANDROID_HOME) $(FYNE) package -os android/arm64

android-install: android
	$(ADB) install -r $(APK)
