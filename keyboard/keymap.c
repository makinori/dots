#include <stdint.h>
#include "keycodes.h"
#include QMK_KEYBOARD_H

enum custom_keycodes {      
    IJKL = SAFE_RANGE,
    COLEMAK,
    RMB_DOWN,
    RMB_UP,
};

#define MOD_LAYER 2
#define MO_MOD MO(MOD_LAYER)

#define MUHENKAN KC_INT5

const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    // qwerty
    [0] = LAYOUT_60_ansi(
        KC_ESC,  KC_1,    KC_2,    KC_3,    KC_4,     KC_5,   KC_6,    KC_7,    KC_8,    KC_9,    KC_0,     KC_MINS, KC_EQL,   KC_BSPC,
        KC_TAB,  KC_Q,    KC_W,    KC_E,    KC_R,     KC_T,   KC_Y,    KC_U,    KC_I,    KC_O,    KC_P,     KC_LBRC, KC_RBRC,  KC_BSLS,
        MO_MOD,  KC_A,    KC_S,    KC_D,    KC_F,     KC_G,   KC_H,    KC_J,    KC_K,    KC_L,    KC_SCLN,  KC_QUOT,           KC_ENT,
        KC_LSFT, KC_Z,    KC_X,    KC_C,    KC_V,     KC_B,   KC_N,    KC_M,    KC_COMM, KC_DOT,  KC_SLSH,                     KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC,                                                        KC_RALT,  MO_MOD,  MUHENKAN, KC_RCTL
    ),
    // colemak dh
    // TODO: maybe not use dh
    [1] = LAYOUT_60_ansi(
        KC_ESC,  KC_1,    KC_2,    KC_3,    KC_4,     KC_5,   KC_6,    KC_7,    KC_8,    KC_9,    KC_0,     KC_MINS, KC_EQL,   KC_BSPC,
        KC_TAB,  KC_Q,    KC_W,    KC_F,    KC_P,     KC_B,   KC_J,    KC_L,    KC_U,    KC_Y,    KC_SCLN,  KC_LBRC, KC_RBRC,  KC_BSLS,
        MO_MOD,  KC_A,    KC_R,    KC_S,    KC_T,     KC_G,   KC_M,    KC_N,    KC_E,    KC_I,    KC_O,     KC_QUOT,           KC_ENT,
        KC_LSFT, KC_X,    KC_C,    KC_D,    KC_V,     KC_Z,   KC_K,    KC_H,    KC_COMM, KC_DOT,  KC_SLSH,                     KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC,                                                        KC_RALT,  MO_MOD,  MUHENKAN, KC_RCTL
    ),
    // all the others keys are here
    [MOD_LAYER] = LAYOUT_60_ansi(
        KC_GRV,  KC_F1,   KC_F2,   KC_F3,   KC_F4,    KC_F5,  KC_F6,   KC_F7,   KC_F8,   KC_F9,   KC_F10,   KC_F11,  KC_F12,   KC_DEL,
        KC_TAB,  KC_MPRV, KC_MPLY, KC_MNXT, KC_NO,    KC_NO,  KC_NO,   KC_PGDN, KC_UP,   KC_PGUP, KC_PSCR,  KC_SCRL, KC_PAUS,  KC_BSLS,
        KC_TRNS, KC_F14,  KC_VOLD, KC_VOLU, KC_MUTE,  KC_NO,  KC_HOME, KC_LEFT, KC_DOWN, KC_RGHT, KC_INS,   KC_DEL,            KC_ENT,
        KC_LSFT, KC_F15,  KC_F16,  KC_F17,  RMB_DOWN, RMB_UP, KC_END,  KC_NO,   KC_COMM, COLEMAK, IJKL,                        KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC,                                                        KC_RALT,  KC_TRNS, MUHENKAN, KC_RCTL
    )
};

bool ijkl_as_arrow_keys = false;

#define HANDLE_AS_USUAL(key)                           \
    if (record->event.pressed) {                       \
        register_code(key);                            \
    } else {                                           \
        unregister_code(key);                          \
    }                                                  \
    return false;

#define IS_QWERTY get_highest_layer(default_layer_state) == 0

bool process_record_user(uint16_t keycode, keyrecord_t *record) {
    switch (keycode) {
        case KC_BSLS:
            if (
                record->event.pressed &&
                layer_state_is(MOD_LAYER) &&
                get_mods() == MOD_BIT(KC_LEFT_CTRL)
            ) {
                reset_keyboard(); // jump to bootloader
                return false;
            } else {
                return true;
            }

        case IJKL:
            if (record->event.pressed) {
                ijkl_as_arrow_keys = !ijkl_as_arrow_keys;
            }
            return false;

        case COLEMAK:
            if (record->event.pressed) {
                set_single_persistent_default_layer(IS_QWERTY ? 1 : 0);
            }
            return false;

        case RMB_DOWN:
            register_code(KC_MS_BTN2);
            return false;

        case RMB_UP:
            unregister_code(KC_MS_BTN2);
            return false;
    }

    if (!ijkl_as_arrow_keys) {
        return true;
    }

    // handle ijkl keys

    if (IS_QWERTY) {
        switch (keycode) {
            case KC_I:
                HANDLE_AS_USUAL(KC_UP);
            case KC_J:
                HANDLE_AS_USUAL(KC_LEFT);
            case KC_K:
                HANDLE_AS_USUAL(KC_DOWN);
            case KC_L:
                HANDLE_AS_USUAL(KC_RIGHT);
        }
    } else {
        switch (keycode) {
            case KC_U:
                HANDLE_AS_USUAL(KC_UP);
            case KC_N:
                HANDLE_AS_USUAL(KC_LEFT);
            case KC_E:
                HANDLE_AS_USUAL(KC_DOWN);
            case KC_I:
                HANDLE_AS_USUAL(KC_RIGHT);
        }
    } 

    return true;
}

#if defined(ENCODER_ENABLE) && defined(ENCODER_MAP_ENABLE)
const uint16_t PROGMEM encoder_map[][NUM_ENCODERS][NUM_DIRECTIONS] = {

};
#endif // defined(ENCODER_ENABLE) && defined(ENCODER_MAP_ENABLE)

