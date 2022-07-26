# Mix those code
# https://github.com/sullivanmatt/json-ups/blob/master/main.py
# https://github.com/vkorobov/ups-telegraf/blob/master/getUpsData.py

from sh import upsc 

import json
import sys

# only needed measurement
measurements=["battery.charge", "battery.voltage", "battery.type", "battery.runtime",
              "input.voltage",
              "ups.load", "ups.beeper.status", "ups.mfr", "ups.model", "ups.serial", "ups.status", "ups.test.result"]

def main():
    # if len(sys.argv) == 2:
        # device = sys.argv[1]
    # else:
        # print("Usage: main.py <UPS device name (e.g 'apc-1500@localhost')>", file=sys.stderr)
        # return

    data = {}
    # $ upsc ups
    status = upsc("ups")
    # print(status)
    # ups.timer.start: 0

    for item in status:
        key = item[:item.find(":")]
        if key in measurements:
            first, second = item.split(': ')
            t = data
            keys = first.split('.')
            second = second.strip()
            for index, key in enumerate(keys):
                if index == len(keys) - 1:
                    t = t.setdefault(key, {'value': second})
                else:
                    t = t.setdefault(key, {})

    # print(json.dumps(data))
    with open('/var/services/web/ups_data.json', 'w') as f_save:
        print(json.dumps(data), file=f_save)

if __name__ == "__main__":
    main()
