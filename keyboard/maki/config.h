// https://docs.qmk.fm/config_options

// https://keebsforall.com/blogs/mechanical-keyboards-101/reduce-keyboard-input-lag-with-qmk

// default is 5
#define DEBOUNCE 3

// https://docs.qmk.fm/feature_debounce_type
// eager down, deferred up, per-key
#define DEBOUNCE_TYPE asym_eager_defer_pk

// 1000 Hz polling. although code shows it's 1 already
#define USB_POLLING_INTERVAL_MS 1 // default 10 (no?)

