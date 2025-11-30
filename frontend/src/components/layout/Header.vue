<template lang="">
    <div class="header_wrapper flex items-center px-4">
        <div class="logo">
            <h1>Perfume App</h1>
        </div>
        <nav class="ml-auto mr-4">
            <ul class="flex space-x-4">
                <li
                    v-for="route in navRoutes"
                    :key="route.path"
                    class="nav-item hover:bg-gray-200 px-2 py-1 rounded"
                >
                    <router-link :to="route.path">
                        <font-awesome-icon
                            v-if="route.meta?.icon"
                            :icon="route.meta.icon"
                        />
                        {{ route.meta.label }}
                    </router-link>
                </li>
            </ul>
        </nav>
        <div>
            <n-dropdown
                v-if="isLoggedIn"
                trigger="hover"
                :options="options"
                @select="handleSelect"
            >
                <div
                    class="rounded-full overflow-hidden w-8 h-8 cursor-pointer"
                >
                    <img
                        src="https://i.pravatar.cc/300"
                        alt="User Avatar"
                        class="w-full h-full object-cover"
                    />
                </div>
            </n-dropdown>
            <n-button v-else type="primary" @click="$router.push('/login')">
                Login
            </n-button>
        </div>
    </div>
</template>
<script>
import { useRouter } from "vue-router";
import { NDropdown, NButton } from "naive-ui";
import { h } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { supabase } from "@/auth/supabase";

function renderIcon(icon) {
    return () => h(FontAwesomeIcon, { icon });
}

export default {
    components: {
        NDropdown,
        NButton,
    },
    data() {
        return {
            session: null,
        };
    },
    setup() {
        const router = useRouter();
        const navRoutes = router.options.routes.filter((r) => r.meta?.nav);

        return {
            navRoutes,
        };
    },
    async created() {
        const { data } = await supabase.auth.getSession();
        this.session = data.session;

        supabase.auth.onAuthStateChange((_event, session) => {
            this.session = session;
        });
    },
    computed: {
        isLoggedIn() {
            return !!this.session;
        },
        options() {
            return [
                { label: "Profile", key: "profile", icon: renderIcon("user") },
                {
                    label: "Settings",
                    key: "settings",
                    icon: renderIcon("cog"),
                },
                {
                    label: "Logout",
                    key: "logout",
                    icon: renderIcon("sign-out-alt"),
                },
            ];
        },
    },
    methods: {
        handleSelect(option) {
            if (option === "logout") {
                this.handleSignOut();
            } else {
                this.$router.push(`/${option}`);
            }
        },
        async handleSignOut() {
            const { error } = await supabase.auth.signOut();
            if (error) {
                console.error(error.message);
            } else {
                this.$router.push("/");
            }
        },
    },
};
</script>
<style lang="scss">
.header_wrapper {
    height: 40px;
}
</style>
