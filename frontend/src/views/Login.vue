<template>
    <div class="login-container">
        <n-card title="Sign In" size="huge" class="login-card">
            <n-form @submit.prevent="handleEmailLogin">
                <n-form-item label="Email">
                    <n-input
                        v-model:value="email"
                        type="email"
                        placeholder="you@example.com"
                    />
                </n-form-item>

                <n-form-item label="Password">
                    <n-input
                        v-model:value="password"
                        type="password"
                        placeholder="••••••••"
                    />
                </n-form-item>

                <n-button
                    type="primary"
                    block
                    :loading="loadingEmail"
                    attr-type="submit"
                >
                    Sign In
                </n-button>
                <div class="footer-text">
                    Don’t have an account?
                    <router-link to="/signup">Register now</router-link>
                </div>
            </n-form>

            <div class="divider">or</div>

            <n-button
                block
                type="info"
                ghost
                :loading="loadingGoogle"
                @click="loginWithGoogle"
            >
                Continue with Google
            </n-button>

            <n-alert v-if="errorMessage" type="error" class="error-alert">
                {{ errorMessage }}
            </n-alert>
        </n-card>
    </div>
</template>

<script>
import { NCard, NForm, NFormItem, NInput, NButton, NAlert } from "naive-ui";
import { supabase } from "@/auth/supabase";

export default {
    name: "Login",
    components: {
        NCard,
        NForm,
        NFormItem,
        NInput,
        NButton,
        NAlert,
    },

    data() {
        return {
            email: "",
            password: "",
            loadingEmail: false,
            loadingGoogle: false,
            errorMessage: "",
        };
    },

    methods: {
        async handleEmailLogin() {
            this.loadingEmail = true;
            this.errorMessage = "";

            const { error } = await supabase.auth.signInWithPassword({
                email: this.email,
                password: this.password,
            });

            this.loadingEmail = false;

            if (error) {
                this.errorMessage = error.message;
            } else {
                this.$router.push("/");
            }
        },

        async loginWithGoogle() {
            this.loadingGoogle = true;
            this.errorMessage = "";

            const { error } = await supabase.auth.signInWithOAuth({
                provider: "google",
                options: {
                    redirectTo: window.location.origin + "/auth/callback",
                },
            });

            this.loadingGoogle = false;

            if (error) {
                this.errorMessage = error.message;
            }
        },
    },
};
</script>

<style scoped>
.login-container {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
}

.login-card {
    width: 360px;
}

.divider {
    text-align: center;
    margin: 16px 0;
    opacity: 0.6;
}

.error-alert {
    margin-top: 16px;
}
</style>
