<pre>
         __              ___    __       
|  |\/| / _` | |\ |  /\   |  | /  \ |\ | 
|  |  | \__> | | \| /~~\  |  | \__/ | \| 
                                         
</pre>

# imgination

Command-line image management utilities. Recursively scans directories
looking for images and reporting on them.

## Features

### Duplicate Image Detection

Finds duplicate images.

#### Usage
`imgination --dir=/Users/crash/Pictures --func=dup`

### GPS Co-ordinate Detection

Finds images with GPS longitude and latitude in the EXIF data.

#### Usage
`imgination --dir=/Users/crash/Pictures --func=gps`

### Minimum Dimension Detection

Finds images with dimensions smaller than specified.

#### Usage
```
imgination --dir=/Users/crash/Pictures --func=dim --width=640
imgination --dir=/Users/crash/Pictures --func=dim --height=480
imgination --dir=/Users/crash/Pictures --func=dim --width=1000 --height=800
```
