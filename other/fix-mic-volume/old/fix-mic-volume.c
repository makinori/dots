#include <wp/wp.h>
#include <glib/gmain.h>
#include <gobject/gobject.h>
#include <gobject/gsignal.h>
#include <stdio.h>
#include <unistd.h>

/*

UNFINISHED AND DOESNT WORK WHEN TRYING TO USE A TIMER

gcc fix-mic-vol-wp.c \
-I /usr/include/wireplumber-0.5 \
-I /usr/include/glib-2.0 \
-I /usr/lib/glib-2.0/include \
-lwireplumber-0.5 -lglib-2.0 -lgobject-2.0 \
-o fix-mic-volume 

*/

// use pw-dump and look for an Audio/Source
#define MIC_NAME "Yeti Nano"
#define MIC_VOLUME 0.7f // %
#define MIC_CHECK_INTERVAL 500 // ms

typedef struct {
	GMainLoop* loop;
	WpCore* core;
	WpObjectManager* om;
	WpPlugin* mixerApi;
	gint exitCode;
} App;

static void on_plugin_loaded(WpCore* core, GAsyncResult* res, App* app)
{
	GError *error = NULL;

	if (!wp_core_load_component_finish(core, res, &error)) {
		fprintf (stderr, "%s\n", error->message);
		app->exitCode = 1;
		g_main_loop_quit (app->loop);
		return;
	}

	app->mixerApi = wp_plugin_find(core, "mixer-api");
	g_object_set(app->mixerApi, "scale", 1 /* cubic */, NULL);
	wp_core_install_object_manager(app->core, app->om);
}

static gboolean on_mic_update_timer(App* app)
{
	printf("timer\n");

	g_auto(GVariantBuilder) b = G_VARIANT_BUILDER_INIT(G_VARIANT_TYPE_VARDICT);
	g_variant_builder_add(&b, "{sv}", "volume", g_variant_new_double(MIC_VOLUME));
  	GVariant* variant = g_variant_builder_end(&b);

	guint32 id;
	gboolean res = FALSE;

	g_autoptr(WpPipewireObject) proxy = wp_object_manager_lookup(
		app->om, WP_TYPE_GLOBAL_PROXY,
		// WP_CONSTRAINT_TYPE_PW_GLOBAL_PROPERTY, "object.id", "=u", id,
		WP_CONSTRAINT_TYPE_PW_PROPERTY, "media.class", "=s", "Audio/Source",
		WP_CONSTRAINT_TYPE_PW_PROPERTY, "node.nick", "=s", MIC_NAME,
		NULL
	);

	printf("%d\n", app);
	printf("%d\n", proxy);

	if (proxy != NULL) {
		id = wp_proxy_get_bound_id(WP_PROXY(proxy));

		printf("%d\n", id);

		g_signal_emit_by_name(app->mixerApi, "set-volume", id, variant, &res);
		// if (!res) {
		// 	fprintf (stderr, "Node %d does not support volume\n", id);
		// 	return;
		// }
	}

	return TRUE;
}

static void on_om_installed(App* app)
{
	

	

	// int checkInterval = 1000000 * MIC_CHECK_INTERVAL;

	

	// // while (1) {
	// printf("hi\n");

	

	// printf("%d\n",proxy);

	GSource* timeoutSource;

	wp_core_timeout_add(
		app->core, &timeoutSource, MIC_CHECK_INTERVAL, (GSourceFunc)on_mic_update_timer, &app, NULL
	);

	// g_source_remove(a); // leave it running

	// g_timer_new()

	// integer_id = gobject.timeout_add( 180000, callback_func)
	// gobject.source_remove(integer_id)

	// }

}

int main()
{
	wp_init(WP_INIT_ALL);

	App app;
	app.loop = g_main_loop_new(NULL, FALSE);
	app.core = wp_core_new(NULL, NULL, NULL);
	app.om = wp_object_manager_new();
	app.exitCode = 0;

	wp_object_manager_add_interest(app.om, WP_TYPE_NODE, NULL);
	wp_object_manager_add_interest(app.om, WP_TYPE_CLIENT, NULL);
	wp_object_manager_request_object_features(
		app.om, WP_TYPE_GLOBAL_PROXY, WP_PIPEWIRE_OBJECT_FEATURES_MINIMAL
	);

	wp_core_load_component(
		app.core, "libwireplumber-module-mixer-api", "module",
		NULL, NULL, NULL, (GAsyncReadyCallback)on_plugin_loaded, &app
	);

	if (!wp_core_connect(app.core)) {
		fprintf (stderr, "Could not connect to PipeWire\n");
		return 2;
	}

	g_signal_connect_swapped(app.core, "disconnected", (GCallback)g_main_loop_quit, app.loop);
	g_signal_connect_swapped(app.om, "installed", (GCallback)on_om_installed, &app);

	g_main_loop_run(app.loop);

	return app.exitCode;
}