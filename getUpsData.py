# Mix those code
# https://github.com/sullivanmatt/json-ups/blob/master/main.py

from sh import upsc 

import json
import sys


def main():
    # if len(sys.argv) == 2:
        # device = sys.argv[1]
    # else:
        # print("Usage: main.py <UPS device name (e.g 'apc-1500@localhost')>", file=sys.stderr)
        # return

    data = {}
    # upsc ups
    status = upsc("ups")
    # print(status)
    # ups.timer.start: 0

    for item in status:
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
