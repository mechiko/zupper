package resource

import (
	_ "embed"
)

//go:embed assets\kfh.png
var pngKfh []byte

const PngKfh = "PngKfh"

//go:embed assets\logo.png
var pngLogo []byte

const PngLogo = "PngLogo"

//go:embed assets\request-64.png
var pngRequester []byte

const PngRequester = "PngRequester"

//go:embed assets\circle.svg
var svgCircle []byte

const SvgCircle = "SvgCircle"

//go:embed assets\request.svg
var svgRequest []byte

const SvgRequest = "SvgRequest"

//go:embed assets\users.ico
var icoUsers []byte

const IcoUsers = "IcoUsers"

//go:embed assets\folder.ico
var icoFolder []byte

const IcoFolder = "IcoFolder"

