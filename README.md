# pantilt
golang package for Pimoroni pantilt using periph/i2c

not (yet?) the cleanest of packages, but it works on my raspi 3. The example test1 uses PtDelta to move the camera around.

Note I am still learning Go, just started on periph,
and this is my first time driving the i2c bus directly, so any tips appreciated.

I made some asumptions about the address of the Pimoroni PanTilt head (0x15) and the i2c bus on the pi (1).

More as I learn more, maybe?

RichR
 Update : modified PtServoEnable to allow each servo to be enabled seperately
 
 (I haven't yet tested only using one at a time... but the correct bit is set.)
 
