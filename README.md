# pag
project architecture generator

## HOW TO SETUP
### Customize JSON.
- please wrire `config/application.json`

#### example
```
{↲
   "goProject":{↲
      "Makefile":"templates/tools/build/Makefile",↲
      "README.md":"templates/documents/readme/README.md",↲
      "config": {↲
          "application.json": ""↲
      },↲
      "cli":{↲
         "${ProjectName}":{↲
            "main.go": "templates/lang/golang/main/hello.go"↲
         }↲
      },↲
      "internal":{↲
         ".keep":""↲
      }↲
   }↲
}↲
```
### Install
```
./install.sh
echo "export PATH=${PATH}:/usr/local/bin" >> ~/.bashrc
```

## HOW TO USE
```
pag add --config #{customizedJSONFilePath}
pag list
pag show --type #{projectType}
pag apply --project-name #{projectName} --output #{path} --type #{projectType}
```
### Example
* Add config file.
``` 
pag add --config config/application.json
``` 

* Project Create.
``` 
pag add --config config/application.json
pag apply --project-name sample_project --output /tmp --type goProject
``` 

* Display Project Types.
``` 
pag list
``` 

* Display Project Tree
``` 
pag show --type goProject
``` 
