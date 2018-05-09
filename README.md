# BIOSCHEMAS.ORG GOWEB

The mapping information on the [Bioschemas.org specification tables](http://bioschemas.org/specifications/) is stored (and transformed into HTML for display purposes) in a human friendly but machine readable format [YAML](http://yaml.org/). A simple way to generate and mantain this mapping is using Google Sheet application given its user experience and user interface which is familiar to anyone who has used Spread sheet aplications before. Thus, we have created a [template](https://docs.google.com/spreadsheets/d/1kl92O05-_3kjYd37YK8q2eb4A1fpYvn3Mkk6HhtUBEs/edit#gid=1483018794) that aids the specification mapping process and provides a suitable way to deliver the information to the website. For more information about the [template and the specification process](https://github.com/BioSchemas/specifications/wiki/Bioschemas-Specification-Process). 

This tool helps you transforms the _Bioschemas Fields_ sheet into a YAML file that can be stored and displayed on the Bioschemas.org specification page. 

## Gettin the specifications CSV
---
Stored on the BIOSCHEMAS.ORG gdrive [specification folder](https://drive.google.com/drive/u/1/folders/0Bw_p-HKWUjHoNThZOWNKbGhOODg) you will find the specification mappings
there on the mapping sheet:
- Go to File > Publish to the web... > Link 
- Select "Bioschemas fields" and "Comma separated values (.csv)" as the output format.
- Click on Publish

If the Publish button is not enable, use the URL provided there. You should get an URL similar to this: 

```https://docs.google.com/spreadsheets/d/e/2PACX-1vTjBKpODC7xTXlfw3P_VffN0u5zdoPt-2bZ4m3_kMGx16lQlz4XH9LyT5y_m-_Tk2AlAO2ID_MFXlTW/pub?gid=292464567&single=true&output=csv```

Use it directly on the tool, or download it and feed the CSV file to the tool using the -f command.

## Anyone can use it!
---
It's just as simple as downloading the right executable file according your Operating System and Architecture.

SO |  Link
--- | ---:
![alt text](images/windows.png "Windows Logo") | [Windows](build/windows)
![alt text](images/ubuntu.png "Linux Logo") | [Linux](build/linux)
![alt text](images/apple.png "Mac Logo") | [MAC Darwin](build/darwin)

> **Note:** Log and YAML result files will be placed on the folder the program is executed.

### Example getting markdown from the GDrive spreadsheet URL
---
#### MAC Darwin / Linux

```./bioschema-goweb -u "https://docs.google.com/spreadsheets/d/e/2PACX-1vTjBKpODC7xTXlfw3P_VffN0u5zdoPt-2bZ4m3_kMGx16lQlz4XH9LyT5y_m-_Tk2AlAO2ID_MFXlTW/pub?gid=292464567&single=true&output=csv"```

#### Windows

```bioschema-goweb_windows_64.exe -u "https://docs.google.com/spreadsheets/d/e/2PACX-1vTjBKpODC7xTXlfw3P_VffN0u5zdoPt-2bZ4m3_kMGx16lQlz4XH9LyT5y_m-_Tk2AlAO2ID_MFXlTW/pub?gid=292464567&single=true&output=csv"```

> **Note:** Use quotation marks for the URL.

### Example getting markdown from the GDrive spreadsheet CSV dowloaded file 

```./bioschema-goweb -f "Beacon Mapping - Bioschemas fields.csv"```

## Help

Use the -h parameter to get info about the command tool.

```./bioschema-goweb -h```
