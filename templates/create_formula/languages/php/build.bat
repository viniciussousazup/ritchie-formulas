:: Java parameters
echo off
SETLOCAL
SET BIN_FOLDER=bin
SET BAT_FILE=%BIN_FOLDER%\run.bat
:build
    rmdir /Q /S %BIN_FOLDER%
    mkdir %BIN_FOLDER%
    xcopy /E /I src %BIN_FOLDER%
    GOTO BAT_WINDOWS
    GOTO CP_DOCKER
    GOTO DONE

:BAT_WINDOWS
    echo start /B /WAIT php -f ./index.php > %BAT_FILE%

:CP_DOCKER
    copy Dockerfile %BIN_FOLDER%
    copy set_umask.sh %BIN_FOLDER%
    GOTO DONE

:DONE
