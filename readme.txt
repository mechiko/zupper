RisohEditor https://github.com/katahiromz/RisohEditor/releases
https://github.com/tc-hib/go-winres

Installation
To install the go-winres command, run:

go install github.com/tc-hib/go-winres@latest
Usage
Please type go-winres help to get a list of commands and options.

Typical usage would be:

Run go-winres init to create a winres directory
Modify the contents of winres.json
Before go build, run go-winres make
go-winres make creates files named rsrc_windows_*.syso that go build automatically embeds in the executable.

The suffix _windows_amd64 is very important. Thanks to it, go build knows it should not include that object in a Linux or 386 build.

Automatic version from git
The --file-version and --product-version flags can take a special value: git-tag. This will retrieve the current tag with git describe --tags and add it to the file properties of the executable.

Using go generate
You can use a //go:generate comment as well:

//go:generate go-winres make --product-version=git-tag
Subcommands
There are other subcommands:

go-winres simply is a simpler make that does not rely on a json file.
go-winres extract extracts resources from an exe file or a dll.
go-winres patch replaces resources directly in an exe file or a dll. For example, to enhance a 7z self extracting archive, you may change its icon, and add a manifest to make it look better on high DPI screens.
JSON format
The JSON file follows this hierarchy:

Resource type (e.g. "RT_GROUP_ICON" or "#42" or "MY_TYPE")
Resource name (e.g. "MY_ICON" or "#1")
Language ID (e.g. "0409" for en-US)
Actual resource: a filename or a json structure
Standard resource types can be found there. But please never use RT_ICON or RT_CURSOR. Use RT_GROUP_ICON and RT_GROUP_CURSOR instead.

Icon JSON
{
  "RT_GROUP_ICON": {
    "APP": {
      "0000": [
        "icon_64.png",
        "icon_48.png",
        "icon_32.png",
        "icon_16.png"
      ]
    },
    "OTHER": {
      "0000": "icon.png"
    },
    "#42": {
      "0409": "icon_EN.ico",
      "040C": "icon_FR.ico"
    }
  }
}
This example contains 3 icons:

"APP"
"OTHER"
42
Windows Explorer will display "APP" because it is the first one. Icons are sorted by name in case sensitive ascending order, then by ID.

42 is an ID, not a name, this is why it comes last.

"APP" is made of 4 png files.
"OTHER" will be generated from one png file. It will be resized to 256x256, 64x64, 48x48, 32x32, and 16x16.
42 is a native icon, it probably already contains several images.
Finally, 42 will display a different icon for french users.

"0409" means en-US, which is the default.
"040C" means fr-FR.
