#include QMK_KEYBOARD_H

enum custom_keycodes {      
    SOCD_W = SAFE_RANGE,
    SOCD_S,
    SOCD_A,
    SOCD_D,

    SOCD_I,
    SOCD_K,
    SOCD_J,
    SOCD_L,

    SOCD_UP,
    SOCD_DOWN,
    SOCD_LEFT,
    SOCD_RIGHT,

    LAYER_0,    
    LAYER_1,   
    LAYER_2,   
    LAYER_3,   
};

#define LAST_LAYER 4

const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    // standard layout
    [0] = LAYOUT_60_ansi(
        KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_MINS, KC_EQL,  KC_BSPC,
        KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_LBRC, KC_RBRC, KC_BSLS,
        MO(LAST_LAYER), KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_QUOT, KC_ENT,
        KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC, KC_RALT, MO(LAST_LAYER), KC_APP, KC_RCTL
    ),
    // standard layout, ijkl are arrow keys
    [1] = LAYOUT_60_ansi(
        KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_MINS, KC_EQL, KC_BSPC,
        KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_UP, KC_O, KC_P, KC_LBRC, KC_RBRC, KC_BSLS,
        MO(LAST_LAYER), KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_LEFT, KC_DOWN, KC_RGHT, KC_SCLN, KC_QUOT, KC_ENT,
        KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC, KC_RALT, MO(LAST_LAYER), KC_APP, KC_RCTL
    ),
    // standard layout, wasd and ijkl are SOCD
    [2] = LAYOUT_60_ansi(
        KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_MINS, KC_EQL,  KC_BSPC,
        KC_TAB, KC_Q, SOCD_W, KC_E, KC_R, KC_T, KC_Y, KC_U, SOCD_I, KC_O, KC_P, KC_LBRC, KC_RBRC, KC_BSLS,
        MO(LAST_LAYER), SOCD_A, SOCD_S, SOCD_D, KC_F, KC_G, KC_H, SOCD_J, SOCD_K, SOCD_L, KC_SCLN, KC_QUOT, KC_ENT,
        KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC, KC_RALT, MO(LAST_LAYER), KC_APP, KC_RCTL
    ),
    // standard layout, ijkl are arrow keys with SOCD
    [3] = LAYOUT_60_ansi(
        KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_MINS, KC_EQL,  KC_BSPC,
        KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, SOCD_UP, KC_O, KC_P, KC_LBRC, KC_RBRC, KC_BSLS,
        MO(LAST_LAYER), KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, SOCD_LEFT, SOCD_DOWN, SOCD_RIGHT, KC_SCLN, KC_QUOT, KC_ENT,
        KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC, KC_RALT, MO(LAST_LAYER), KC_APP, KC_RCTL
    ),
    // all the others keys are here
    [LAST_LAYER] = LAYOUT_60_ansi(
        KC_GRV, KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, KC_DEL,
        KC_TAB, KC_MPRV, KC_MPLY, KC_MNXT, KC_R, KC_T, KC_Y, KC_PGUP, KC_UP, KC_PGDN, KC_PSCR, KC_SCRL, KC_PAUS, KC_BSLS,
        KC_TRNS, KC_F14, KC_VOLD, KC_VOLU, KC_MUTE, KC_G, KC_HOME, KC_LEFT, KC_DOWN, KC_RGHT, KC_INS, KC_DEL, KC_ENT,
        KC_LSFT, KC_F15, KC_F16, KC_F17, KC_V, KC_B, KC_END, LAYER_3, LAYER_2, LAYER_1, LAYER_0, KC_RSFT,
        KC_LCTL, KC_LGUI, KC_LALT, KC_SPC, KC_RALT, KC_TRNS, KC_APP, KC_RCTL
    )
};

bool socd_KC_W = false;
bool socd_KC_S = false;
bool socd_KC_A = false;
bool socd_KC_D = false;

bool socd_KC_I = false;
bool socd_KC_K = false;
bool socd_KC_J = false;
bool socd_KC_L = false;

bool socd_KC_UP = false;
bool socd_KC_DOWN = false;
bool socd_KC_LEFT = false;
bool socd_KC_RIGHT = false;

#define SOCD_KEY(a, b)                                \
    if (record->event.pressed) {                      \
        socd_##a = true;                              \
        if (socd_##b) unregister_code(b);             \
        register_code(a);                             \
    } else {                                          \
        socd_##a = false;                             \
        unregister_code(a);                           \
        if (socd_##b) register_code(b);               \
    }                                                 \
    return false;


#define LAYER_KEY(i)                                  \
    if (record->event.pressed) {                      \
        set_single_persistent_default_layer(i);       \
        register_code(KC_F2##i);                      \
    } else {                                          \
        unregister_code(KC_F2##i);                    \
    }                                                 \
    return false;


bool process_record_user(uint16_t keycode, keyrecord_t *record) {
    switch (keycode) {
        case SOCD_W:
            SOCD_KEY(KC_W, KC_S)
        case SOCD_S:
            SOCD_KEY(KC_S, KC_W)
        case SOCD_A:
            SOCD_KEY(KC_A, KC_D)
        case SOCD_D:
            SOCD_KEY(KC_D, KC_A)

        case SOCD_I:
            SOCD_KEY(KC_I, KC_K)
        case SOCD_K:
            SOCD_KEY(KC_K, KC_I)
        case SOCD_J:
            SOCD_KEY(KC_J, KC_L)
        case SOCD_L:
            SOCD_KEY(KC_L, KC_J)

        case SOCD_UP:
            SOCD_KEY(KC_UP, KC_DOWN)
        case SOCD_DOWN:
            SOCD_KEY(KC_DOWN, KC_UP)
        case SOCD_LEFT:
            SOCD_KEY(KC_LEFT, KC_RIGHT)
        case SOCD_RIGHT:
            SOCD_KEY(KC_RIGHT, KC_LEFT)

        case LAYER_0:
            LAYER_KEY(0)
        case LAYER_1:
            LAYER_KEY(1)
        case LAYER_2:
            LAYER_KEY(2)
        case LAYER_3:
            LAYER_KEY(3)

        case KC_BSLS:
            if (
                record->event.pressed &&
                layer_state_is(LAST_LAYER) &&
                get_mods() == MOD_BIT(KC_LEFT_CTRL)
            ) {
                reset_keyboard(); // jump to bootloader
                return false;
            } else {
                return true;
            }

        default:
            return true;
    }
}

#if defined(ENCODER_ENABLE) && defined(ENCODER_MAP_ENABLE)
const uint16_t PROGMEM encoder_map[][NUM_ENCODERS][NUM_DIRECTIONS] = {

};
#endif // defined(ENCODER_ENABLE) && defined(ENCODER_MAP_ENABLE)

