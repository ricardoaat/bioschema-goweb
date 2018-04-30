# BIOSCHEMAS.ORG GOWEB

Transforms specifications definitions on the bioschemas.org's drive repository into ready-to-include Jekyll-Markdown files for [Bioschemas.org](http://bioschemas.org)

## Gettin the specifications CSV
Stored on the BIOSCHEMAS.ORG gdrive [specification folder](https://drive.google.com/drive/u/1/folders/0Bw_p-HKWUjHoNThZOWNKbGhOODg) you will find the specification mappings
there on the mapping sheet:
- Go to File > Publish to the web... > Link 
- Select "Bioschemas fields" and "Comma separated values (.csv)" as the output format.
- Click on Publish

If the Publish button is not enable, use the URL provided there. You should get an URL similar to this: 

``` https://docs.google.com/spreadsheets/d/e/2PACX-1vRb23ofqiLsw4AvBFr5ysBdrtNvsGCYc4CdPhPPaeLd0MJm3mhCOYSdidlwhkGi6mqEJ7Qx9rLQcCPq/pub?gid=292464567&single=true&output=csv ```

Use it directly on the tool, or download it and feed the CSV file to tool using the -f command.

## Anyone can use it!

It's just as simple as downloading the right executable file according your Operating System and Architecture.

SO |  Link
--- | ---:
![alt text](images/windows.png "Windows Logo") | [Windows](build/windows)
![alt text](images/ubuntu.png "Linux Logo") | [Linux](build/linux)
![alt text](images/apple.png "Mac Logo") | [MAC Darwin](build/darwin)

Once you have got the program, make sure to have a log folder right next to it, as well as the [config.toml](build/example/config.toml) file. On this file you can configure the path in which the logs will be stored. 

### Example getting markdown from the GDrive spreadsheet URL

```./bioschema-goweb -u "https://docs.google.com/spreadsheets/d/e/2PACX-1vTgQwURl7JI0S_ZcbgMrvkGZhQDLJv322vb86Vbdt1k0_AbIpoS_fAebeJGG0U328Kewqtm_3Ww2YDK/pub\?gid\=292464567\&single\=true\&output\=csv"```

> **Note:** Use quotation marks for the URL

### Example getting markdown from the GDrive spreadsheet CSV file 

```./bioschema-goweb -f "Beacon Mapping - Bioschemas fields.csv"```

## Help

Use the -h parameter to get info about the command tool.

```./bioschema-goweb -h```
