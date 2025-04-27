# Extras for performance

-   refresh your mirrorlist for fastest

    `sudo reflector --country US --latest 25 --score 25 --sort rate --protocol https --verbose --save /etc/pacman.d/mirrorlist`

-   set cpu frequency policy to performance

    -   install cpupower `sudo pacman -S cpupower`
    -   check available governors and current policy `cpupower frequency-info`
    -   modify `/etc/default/cpupower` and set governer to `performance`
    -   enable and start systemd service `sudo systemctl enable --now cpupower`
    -   check if policy was updated `cpupower frequency-info`

    find more info here https://wiki.archlinux.org/title/CPU_frequency_scaling#cpupower

-   TODO: [linux-tkg](https://github.com/Frogging-Family/linux-tkg)

-   improve performance for gaming

    -   https://wiki.archlinux.org/title/Gaming#Increase_vm.max_map_count

        `sudo vim /etc/sysctl.d/80-gamecompatibility.conf`

        ```
        vm.max_map_count = 2147483642
        ```

        `sudo sysctl --system` to apply

    -   https://wiki.archlinux.org/title/Gaming#Tweaking_kernel_parameters_for_response_time_consistency

        `sudo vim /etc/tmpfiles.d/consistent-response-time-for-gaming.conf`

        ```
        #    Path                  Mode UID  GID  Age Argument
        w /proc/sys/vm/compaction_proactiveness - - - - 0
        w /proc/sys/vm/watermark_boost_factor - - - - 1
        w /proc/sys/vm/min_free_kbytes - - - - 1048576
        w /proc/sys/vm/watermark_scale_factor - - - - 500
        w /proc/sys/vm/swappiness - - - - 10
        w /sys/kernel/mm/lru_gen/enabled - - - - 5
        w /proc/sys/vm/zone_reclaim_mode - - - - 0
        w /sys/kernel/mm/transparent_hugepage/enabled - - - - madvise
        w /sys/kernel/mm/transparent_hugepage/shmem_enabled - - - - advise
        w /sys/kernel/mm/transparent_hugepage/defrag - - - - never
        w /proc/sys/vm/page_lock_unfairness - - - - 1
        w /proc/sys/kernel/sched_child_runs_first - - - - 0
        w /proc/sys/kernel/sched_autogroup_enabled - - - - 1
        w /proc/sys/kernel/sched_cfs_bandwidth_slice_us - - - - 3000
        w /sys/kernel/debug/sched/base_slice_ns  - - - - 3000000
        w /sys/kernel/debug/sched/migration_cost_ns - - - - 500000
        w /sys/kernel/debug/sched/nr_migrate - - - - 8
        ```

-   improve startup time

    check using `systemd-analyze critical-chain`

    -   `ldconfig.service` may take a while if lots of packages are installed

        disable with `sudo systemctl mask ldconfig.service`
