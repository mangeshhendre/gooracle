# gooracle

Connecting to Oracle
--------------------

### Set up for windows:
1. Install Go 32 bit version

2. Install Oracle instant client 32bit version
    * http://www.oracle.com/technetwork/topics/winsoft-085727.html
    * Create folder C:\instantclient_12_2_32
    * Extract instant client basic to C:\instantclient_12_2_32
    * Extract instant client sdk to  C:\instantclient_12_2_32\sdk
    * Add C:\instantclient_12_2_32 to the path environment variable
    
3. Install TDM-GCC-32
    * http://tdm-gcc.tdragon.net/download (tdm-gcc-5.1.0-3.exe)
    * Add C:\TDM-GCC-32\bin to the 'path' environment variable
    
4. pkg-config setup
    * Create C:\pkg-config folder
    * Extract the folowing to C:\pkg-config
    
      http://ftp.gnome.org/pub/gnome/binaries/win32/dependencies/pkg-config_0.26-1_win32.zip
      
      http://ftp.gnome.org/pub/gnome/binaries/win32/glib/2.28/glib_2.28.8-1_win32.zip
      
      http://ftp.gnome.org/pub/gnome/binaries/win32/dependencies/gettext-runtime_0.18.1.1-2_win32.zip
      
    * Copy C:\pkg-config\glib_2.28.8-1_win32\bin\libglib-2.0-0.dll to C:\pkg-config\pkg-config_0.26-1_win32\bin
    * Copy C:\pkg-config\gettext-runtime_0.18.1.1-2_win32\bin\intl.dll to C:\pkg-config\pkg-config_0.26-1_win32\bin
    * Add C:\pkg-config\pkg-config_0.26-1_win32\bin to the 'path' environment variable
   
 5. Create the oci8.pc file
     * Add the oci8.pc file to C:\pkg-config\config
     * Set PKG_CONFIG_PATH environemnt variable to C:\pkg-config\config
    
    ### oci8.pc file
    
    ```
    libdir=C:/instantclient_12_2_32/sdk/lib/msvc
    includedir=C:/instantclient_12_2_32/sdk/include

    glib_genmarshal=glib-genmarshal
    gobject_query=gobject-query
    glib_mkenums=glib-mkenums

    Name: oci8
    Description: oci8 library
    Libs: -L${libdir} -loci
    Cflags: -I${includedir}
    Version: 12.2
    ```
Install Oracle Driver
---------------------
    go get github.com/mattn/go-oci8
        
