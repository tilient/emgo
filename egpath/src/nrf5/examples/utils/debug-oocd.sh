#!/bin/sh

set -e

arch=`grep 'EGARCH=' ../build.sh |sed 's/.*EGARCH=\([[:alnum:]_]\+\).*/\1/g'`
if [ -z "$arch" ]; then
	arch=$EGARCH
fi

oocd_cmd="openocd -f interface/$INTERFACE.cfg -f target/$TARGET.cfg -c 'gdb_port pipe' -c 'log_output /dev/null'"

arm-none-eabi-gdb --tui \
	-ex "target extended-remote | $oocd_cmd" \
	-ex 'set mem inaccessible-by-default off' \
	-ex "set remote hardware-breakpoint-limit 4" \
	-ex "set remote hardware-watchpoint-limit 2" \
	-ex 'set history save on' \
	-ex 'set history filename ~/.gdb-history-emgo' \
	-ex 'set history size 1000' \
	-ex 'monitor reset init' \
	$arch.elf
