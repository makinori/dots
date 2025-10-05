import weechat

SCRIPT_NAME = "indent_tox"
SCRIPT_AUTHOR = "maki"
SCRIPT_VERSION = "0.1"
SCRIPT_LICENSE = "GPL3"
SCRIPT_DESC = "Indent Tox"

# use /buffer localvar to get info

def indent_buffer(buffer_ptr):
	plugin = weechat.buffer_get_string(buffer_ptr, "plugin")
	name = weechat.buffer_get_string(buffer_ptr, "name")
	if plugin and plugin == "tox":
		if name and "/" in name:
			weechat.buffer_set(buffer_ptr, "localvar_set_type", "private")
		else:
			weechat.buffer_set(buffer_ptr, "localvar_set_type", "server")
			weechat.buffer_set(buffer_ptr, "short_name", "tox:" + name)

def buffer_opened_cb(data, signal, signal_data):
	indent_buffer(signal_data)
	return weechat.WEECHAT_RC_OK

def indent_current():
	infolist = weechat.infolist_get("buffer", "", "")
	if infolist == None:
		return
	while weechat.infolist_next(infolist):
		buffer_ptr = weechat.infolist_pointer(infolist, "pointer")
		if buffer_ptr == None:
			return
		indent_buffer(buffer_ptr)

if weechat.register(
    SCRIPT_NAME, SCRIPT_AUTHOR, SCRIPT_VERSION, SCRIPT_LICENSE, SCRIPT_DESC, "",
    ""
):
	indent_current()
	weechat.hook_signal("buffer_opened", "buffer_opened_cb", "")
