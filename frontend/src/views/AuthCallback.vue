<template>
    <div class="auth-callback">
        <n-spin size="large">Logging you in…</n-spin>
    </div>
</template>

<script>
import { NSpin } from "naive-ui";
import { supabase } from "@/auth/supabase";

export default {
    name: "AuthCallback",

    components: {
        NSpin,
    },

    data() {
        return {};
    },

    mounted() {
        this.handleAuthCallback();
    },

    methods: {
        async handleAuthCallback() {
            try {
                // Supabase automatically completes the OAuth process
                const { data, error } = await supabase.auth.getSession();

                if (error) {
                    console.error("Error getting session:", error.message);
                }

                if (data.session) {
                    this.$router.push("/"); // redirect to home if logged in
                } else {
                    this.$router.push("/login"); // redirect to login if not
                }
            } catch (err) {
                console.error("Auth callback error:", err);
                this.$router.push("/login");
            }
        },
    },
};
</script>

<style scoped>
.auth-callback {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
