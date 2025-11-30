<template>
    <div class="signup-container">
        <n-card title="Create Account" size="huge" class="signup-card">
            <n-form @submit.prevent="handleSignUp">
                <!-- Email -->
                <n-form-item label="Email">
                    <n-input
                        v-model:value="email"
                        type="email"
                        placeholder="you@example.com"
                    />
                </n-form-item>

                <!-- Password -->
                <n-form-item label="Password">
                    <n-input
                        v-model:value="password"
                        type="password"
                        placeholder="••••••••"
                    />
                </n-form-item>

                <!-- Confirm Password -->
                <n-form-item label="Confirm Password">
                    <n-input
                        v-model:value="confirmPassword"
                        type="password"
                        placeholder="••••••••"
                    />
                </n-form-item>

                <n-button
                    type="primary"
                    block
                    :loading="loading"
                    attr-type="submit"
                >
                    Sign Up
                </n-button>
            </n-form>

            <div class="footer-text">
                Already have an account?
                <router-link to="/login">Login here</router-link>
            </div>

            <n-alert
                v-if="errorMessage"
                type="error"
                class="error-alert"
                closable
            >
                {{ errorMessage }}
            </n-alert>

            <n-alert
                v-if="successMessage"
                type="success"
                class="success-alert"
                closable
            >
                {{ successMessage }}
            </n-alert>
        </n-card>
    </div>
</template>

<script>
import { NCard, NForm, NFormItem, NInput, NButton, NAlert } from "naive-ui";
import { supabase } from "@/auth/supabase";
import { postProfile } from "@/composables/useSupabase";
import { useProfiles } from "@/composables/useProfiles";

export default {
    name: "SignUp",
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
            confirmPassword: "",
            loading: false,
            errorMessage: "",
            successMessage: "",
        };
    },
    methods: {
        validatePassword(password) {
            const rules = [
                { regex: /.{8,}/, message: "At least 8 characters" },
                { regex: /[A-Z]/, message: "At least one uppercase letter" },
                { regex: /[a-z]/, message: "At least one lowercase letter" },
                { regex: /[0-9]/, message: "At least one number" },
                {
                    regex: /[^A-Za-z0-9]/,
                    message: "At least one special character",
                },
            ];

            const failed = rules.filter((rule) => !rule.regex.test(password));
            if (failed.length) {
                return failed.map((f) => f.message).join(", ");
            }
            return null;
        },

        async handleSignUp() {
            this.loading = true;
            this.errorMessage = "";
            this.successMessage = "";

            // 1️⃣ Validate passwords match
            if (this.password !== this.confirmPassword) {
                this.errorMessage = "Passwords do not match.";
                this.loading = false;
                return;
            }

            // 2️⃣ Validate password strength
            const passwordError = this.validatePassword(this.password);
            if (passwordError) {
                this.errorMessage = `Password is not strong enough: ${passwordError}`;
                this.loading = false;
                return;
            }

            try {
                // 3️⃣ Attempt signup
                const { data, error } = await supabase.auth.signUp({
                    email: this.email.toLowerCase(),
                    password: this.password,
                });

                // Check if email is already registered
                const { checkExistingEmail } = useProfiles();
                const emailExists = checkExistingEmail(this.email);
                if (emailExists) {
                    this.errorMessage = "Email is already registered.";
                    this.loading = false;
                    return;
                }

                if (error) {
                    this.errorMessage = error.message;
                    this.loading = false;
                    return;
                }

                // 4️⃣ Signup success
                this.successMessage =
                    "Account created! Please check your email to confirm your account.";

                // 5️⃣ Create user profile in database
                if (data.user?.id) {
                    try {
                        await postProfile({
                            userId: data.user.id,
                            email: this.email,
                        });
                    } catch (profileError) {
                        console.error("Error creating profile:", profileError);
                    }
                }
            } finally {
                this.loading = false;
            }
        },
    },
};
</script>

<style scoped>
.signup-container {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
}

.signup-card {
    width: 360px;
}

.footer-text {
    margin-top: 16px;
    font-size: 14px;
    text-align: center;
}

.error-alert,
.success-alert {
    margin-top: 16px;
}
</style>
