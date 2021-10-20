# OsintDB #

## What is it ##

This project is designed to be an OSINT resource collection manager. It helps keeping all your OSINT resources in one
place and also provides some filtering capabilities, so you can easily find the resource you need for your task.
Currently, it only features a GUI.

## How to use it ##

### View the Tools ###

Simply start the application and browser through the data. By clicking on an entry the Detail View is opened.

### Filter the Resources ###

By clicking the "funnel"-Icon the Filtering Sidebar is opened. Currently, OsintDB supports realtime fuzzy filtering for
the following fields: title, type (such as App, Link or Browser Extension), description, URL, user-defined comment,
tags, world regions, licenses and browsers (for browser extensions).

### Modify Entries ###

Inside the Detail View every information can be edited by first Clicking on the "Edit" button. Don't forget to save your
changes by clicking "Save".

### Delete Entries ###

Inside the Detail View click "Delete". 

### Add Entries ###

Clicking the "+"-Icon lets you create new entries.

### Import Data ####

OsintDB already provides a curated list of OSINT resources out-of-the-box. If you want to import your own data you will
have to prepare it as a JSON file and provide the required metadata for each tool. Then you can easily import the Data
in the Settings menu.

### Export Data ###

All the data in the application can easily be exported from the settings menu.
  
## Installation ##

Note: Before executing osintDB it is recommended to create `~/.local/osintdb` on your machine and copy `tools.min.json`
to it so osintDB can find it and import the tool collection on first startup.

### From Source ###

The following dependencies need to be setup on your machine:

- [Golang](https://golang.org/)
- [Wails](https://wails.app/)  
- [ObjectBox](https://objectbox.io/) (at least `libobjectbox.so`)
- [npm](https://www.npmjs.com/) 

After installing them the project can be built by running `wails build` in the project root directory. The binary can
then be found in the `build` folder.

### Releases ###

Even when using the release versions `libobjectbox.so` needs to be setup on your machine. It can be
found [here](https://github.com/objectbox/objectbox-c) (e.g. just extract the latest release to get it).

The releases need to be extracted locally and then executed.

## License ##

The project was released under the terms and conditions of the MIT License.