# <!--  Home -->

![](Corporate%20Identity/6inarow_temp.png)  
Team 6InARow (CEP 2021 - Hochschule Mannheim)

[![main branche](https://github.com/6inarow/Auditheia/actions/workflows/.github/workflows/gitAutomatedTesting_2.yml/badge.svg)](https://github.com/6inarow/Auditheia/actions/workflows/gitAutomatedTesting_2.yml) 
[![GitHub license](https://img.shields.io/github/license/GhostWriters/DockSTARTer.svg?style=flat-square&color=607D8B)](https://github.com/6inarow/Auditheia/blob/main/LICENSE)

The Framework Auditheia (from "Audit" and "Aletheia", greek goddess, spirit of truth) helps to gather the configurations of a system for further analysis and is compatible with Windows10, Linux and macOS operating systems.

Tested operating Systems:
> Windows 10, Builds 14393 and 17763 \
MacOS 10.15.7, Build 19H1217 \
Linux: Ubuntu 18.04.5 LTS, Ubuntu 20.04.2 LTS


# Usage
On Linux and macOS the program has to be run directly through the command line, on Windows it's also possible to create a shortcut with all needed arguments.
To gather information about your server or PC you need to create a configuration file describing what to test and how.
As result all reports and audit files get stored in the output folder. (default: auditFiles)

## Configuration
The configuration file contains some metadata, a list of files to always gather, and a list of audits.
These audits contain a list of tasks which either are a system command, or a script (ECMAScript 5.1) to execute.
By providing a set of expected values and an appropriate operator, Auditheia is able to automatically compare the results of the command or script with the expected values.
Each task may have one of the three following states:

* "Success" : the tasks result is as expected according to the operator
* "Fail" : the tasks result is not as expected according to the operator
* "Error" : execution of the specified command or script failed

For each of these three states one task can be defined if desired, which will be executed after its parent task entered the specified state.

## Artefacts
While executing defined tasks, the framework scans each command for file paths and saves these as artefacts.  
Example: upon executing the command `cat /etc/ssh/sshd_conf` on a linux system, the framework identifies `/etc/ssh/sshd_conf` as an artefact file and attempts to save it in the frameworks output directory.  
Additionally, all command outputs will be saved as artefacts too.

## Report
After finishing execution of all audits and their respective tasks, a report will be generated and written to the output directory.
This report contains some metadata, as well as the results of each executed task.
If an error occurred during task execution, a status message with further details about the error will be provided.
The report will be either in JSON or YAML format, which can be defined by providing the corresponding flag.
By default, the report will be in JSON format.

## Logs
While the framework is running, all messages of all levels will be written to a log file and the console in parallel.
These messages will always have different levels ranging from "error" to "trace".
Each loglevel always includes all levels below itself. (see [Parameters](#parameters))
As for all outputs the logfile too will be located inside the output directory.
If creation of the logfile fails, log messages will still be visible inside the console.

## Parameters
|Parameter|Description|
|---|---|
| --output|Path to your output folder where all log and report files are stored to|
| --conf|Name of your Config file|
| --json|Report format JSON (default) if both json and yaml are set to true, json will be used|
| --yaml|Report format YAML|
| --report|Name of the Report file|
| --check-only|Sanity check the program will only check if the config is readable and all executables do exist|
| --verbosity|Output depth <br> 0 = Output of all Errors <br> 1 = Output of warnings and all previous levels <br> 2 = Output of infos and all previous levels <br> 3 = Output of debug and all previous levels <br> 4 = Output of trace (output + code file and line) and all previous levels|

example:
./auditheia --conf="audit.json" --verbosity=3

## JSON Config file (descriptive)
```json
{
  "customer_name": "Name of the customer",
  "initial_date": "Date when the config file got created",
  "last_changed": "Date when the last change to the config file was",
  "version": "Version of the config file",
  "conf_os": "What operation system is used at execution (Windows, Linux or MacOS)",
  "root_required": "Do you need root/admin permissions? (true or false)",
  "additional_files":[ 
    "Array of paths/files", 
    "you want to get saved in the report"
  ],
  "audit_list":[
   {
    "name": "Name of the  audit",
    "tasks":{
        "type": "type of the execute (command or script)",
        "execute": "command or script to execute",
        "expected": [
          "results of execute"
        ],
        "operator": "result of expected gets compared with one of the operator : equals, not equals, contains, not contains, lesser, greater",
        "on_success": {
          "type": "type of the execute (command or script)",
          "execute": "command or script to execute",
          "expected": [
            "results of execute"
          ],
          "operator": "result of expected gets compared with one of the operator : equals, not equals, contains, not contains, lesser, greater",
          "on_failure": {
            "type": ".... and so on, no limits"
           }
         },
        "on_failure": {
          "type": "type of the execute (command or script)",
          "execute": "command or script to execute",
          "expected": [
            "results of execute"
           ],
          "operator": "result of expected gets compared with one of the operator : equals, not equals, contains, not contains, lesser, greater"
         },
         "on_error": {
           "type": "type of the execute (command or script)",
           "execute": "command or script to execute",
           "expected": [
            "results of execute"
           ],
           "operator": "result of expected gets compared with one of the operator : equals, not equals, contains, not contains, lesser, greater"
         }
      }
    }
  ]
}
```
## YAML Config file (example)
```yaml
---
customer_name: "Mr. Anderson"
initial_date: '2021-05-06'
last_changed: '2021-05-06'
version: "0.0.1"
conf_os: "linux"
root_required: true
additional_files:
  - "/etc/shadow"
  - "/etc/passwd"
audit_list:
  - name: "cramfs filesystem check"
    tasks:
      - type: "command"
        executable: "modprobe -n -v cramfs"
        expected:
          - "install /bin/true"
        operator: "equals"
        on_success:
          type: "command"
          executable: "lsmod | grep cramfs"
          expected:
            - ''
          operator: "equals"
  - name: "echo test"
    tasks:
      - type: "command"
        executable: "echo hello"
        expected:
          - hello
        operator: "equals"
        on_success:
          type: "command"
          executable: "echo test2"
          expected:
            - "test3"
          operator: "not equals"
  - name: "script test"
    tasks:
      - type: "script"
        executable: "let result;\
        let commandResult = auditheia.osExec('echo hello');\
        result=commandResult[1];\
        console.log('result: ' + result);\
        console.log('commandResult:'+ commandResult);\
        let err = commandResult[0];"
        expected:
          - "hello"
        operator: "equals"
```
# Development

If you want to add more features or have some more improvement of life ideas, please consider making a Pull Request.

## Requirements
Go 1.16.x + \
install external libraries via _go get_ [as of Go 1.17+ _go get_ is deprecated use _go install_ instead](https://golang.org/doc/go-get-install-deprecation)\
[goja](https://github.com/dop251/goja) \
[lumberjack](https://github.com/natefinch/lumberjack) \
[logurs](https://github.com/sirupsen/logrus) \
[yaml](https://gopkg.in/yaml.v2) \
[sys](https://golang.org/x/sys)

IDE: \
[Visual Studio Code](https://code.visualstudio.com/) with the [go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) \
[JetBrains Goland](https://www.jetbrains.com/go/)
