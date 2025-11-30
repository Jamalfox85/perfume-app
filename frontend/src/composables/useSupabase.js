import { supabase } from "@/auth/supabase";
import api from "../services/apiClient";

export async function loginWithPass(email, password) {
    const { data, error } = await supabase.auth.signInWithPassword({
        email,
        password,
    });
    if (error) throw error;
    return data.session; // contains the access token
}

export async function loginWithGoogle() {
    await supabase.auth.signInWithOAuth({
        provider: "google",
        options: {
            redirectTo: "http://localhost:5173/auth/callback",
        },
    });
}

export async function postProfile(profileData) {
    // explode properties
    let data = {
        userId: profileData.userId,
        email: profileData.email,
    };
    try {
        const response = await api.post("/profiles", data);
        return response.data;
    } catch (error) {
        console.error("Error creating profile:", error);
        throw error;
    }
}
