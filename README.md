# pantilt
golang package for Pimoroni pantilt using periph/i2c

Works well on my raspi 3. The example test1 uses Delta to move the camera around.

I made some asumptions about the address of the Pimoroni PanTilt head (0x15) and the i2c bus on the pi (1). These are constants in the package. I changed the max for servoMax compared to the Pimoroni values 2375 - 575 gives a better (1800) range that produces closer to the required angle on my servos.

RichR
 Update : modified ServoEnable to allow each servo to be enabled seperately
 and changed some names... e.g. Home replaced by Go - see the example for more

update: spotted why the servos gave a small jitter at the end of delta travel :- the for statement was one short of the correct
value. see lines 179 / 184 of pantilt.go...
