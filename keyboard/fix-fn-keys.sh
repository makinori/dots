#!/bin/sh

dir=$(mktemp -d)
mkdir "${dir}"/symbols
cat > "${dir}"/symbols/myoptions <<EOF
xkb_symbols "restore_fk" {
    key<FK13> { [ F13 ] };
    key<FK14> { [ F14 ] };
    key<FK15> { [ F15 ] };
    key<FK16> { [ F16 ] };
    key<FK17> { [ F17 ] };
    key<FK18> { [ F18 ] };
    key<FK19> { [ F19 ] };
    key<FK20> { [ F20 ] };
    key<FK21> { [ F21 ] };
    key<FK22> { [ F22 ] };
    key<FK23> { [ F23 ] };
    key<FK24> { [ F24 ] };
};
EOF

setxkbmap -print \
    | sed 's/\(xkb_symbols.*\)"/\1+myoptions(restore_fk)"/' \
    | xkbcomp -I"${dir}" - $DISPLAY

rm -r "${dir}"
