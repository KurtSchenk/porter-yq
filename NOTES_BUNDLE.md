* First create, compile and install yq mixin. See [NOTES_MAGE.md](./NOTES_MAGE.md)
* There is a test bundle that uses this yq mixin. To run go to ./tests/bundle
* And there you can run 

```
bash (base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/tests/bundle# porter install
executing install action from porter-custom-mixin (installation: /porter-custom-mixin)
Install Hello World
Hello World
yq mixin version
yq 3.2.3
Extract name from data.xml using yq
"My name"
Echo the name saved to output from the previous step
My name
execution completed successfully!
```

* Show bundle output from installation

```
bash (base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/tests/bundle# porter installation list
---------------------------------------------------------------------------------
  NAMESPACE  NAME                 VERSION  STATE        STATUS     MODIFIED      
---------------------------------------------------------------------------------
             porter-custom-mixin  0.1.0    installed    succeeded  1 minute ago  
                                  0.1.1    uninstalled  succeeded  2023-11-01    

(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/tests/bundle# porter installations output list
------------------------------------
  Name         Type    Value        
------------------------------------
  name         string  "My name"    
                                    
  test_output  string  Hello World
  ```