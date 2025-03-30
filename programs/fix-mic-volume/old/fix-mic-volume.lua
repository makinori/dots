#!/usr/bin/wpexec

-- TODO: this hangs i think the moment you log in
-- hopefully the added sleep fixes that for now

mic_name = "Yeti Nano"
mic_volume = 0.6 -- %
mic_interval = 500 -- ms

--

mic_interest = Interest {
	type = "node",
	Constraint { "media.class", "matches", "Audio/Source" },
	Constraint { "node.nick", "matches", mic_name }
}

obj_mgr = ObjectManager {
	mic_interest
}

mic_volume = mic_volume ^ 3 -- cubic

Core.require_api("mixer", function(mixer)
	obj_mgr:connect("installed", function(om)
		Core.timeout_add(mic_interval, function()
			local mic = om:lookup(mic_interest)
			if mic ~= nil then
				mixer:call("set-volume", mic["bound-id"], mic_volume)
			end
			return true
		end)
	end)
	obj_mgr:activate()
end)

