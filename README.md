# pi-featherwing

Various Rust-based tools for controlling Adafruit Featherwings from a Raspberry Pi.

## Disclaimer

While the code is copyright Google, LLC, this is not an official Google project - it is just a bunch of weekend experiments that I did on my own.

## Wiring

I used an [Adafruit Perma Proto Bonnet Mini Kit](https://www.adafruit.com/product/3203) and wired it to a [FeatherWing Proto](https://www.adafruit.com/product/2884).

I connected it as follows. In the table, I use the [GPIO/BCM pin names](https://pinout.xyz/) and an extended version of the [Feather pin names](https://learn.adafruit.com/adafruit-feather/feather-specification#pin-naming-3096825-19). The layout of the diagram is per the standard Featherwing physical configuration (looking top down).

| Left                 |             Right |
| :------------------- | ----------------: |
| Rst (NC)             |                   |
| * 3.3V <- +3V        |                   |
| * Aref <- +3V        |                   |
| * GND - GND          |                   |
| A0 <-> GPIO 20       |         (NC) VBAT |
| A1 <-> GPIO 21       |   GPIO 27 -> EN * |
| A2 <-> GPIO 22       |         (NC) VBUS |
| A3 <-> GPIO 23       |   GPIO 13 <-> D13 |
| A4/D24 <-> GPIO 24   |   GPIO 12 <-> D12 |
| A5/D25 <-> GPIO 25   |   GPIO 19 <-> D11 |
| * SCK <- CLK         | GPIO 18 <-> D10 * |
| * MO <- MOSI         |  GPIO 17 <-> D9 * |
| * MI <- MISO         |   GPIO 6 <-> D6 * |
| RX/D0 (NC for now)   |   GPIO 5 <-> D5 * |
| TX/D1 (NC for now)   |      SCL -> SCL * |
| D4 <-> GPIO 4        |     SDA <-> SDA * |

Asterisk (*) indicates connection is implemented.