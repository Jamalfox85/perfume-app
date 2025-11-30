import api from "../services/apiClient";
import { ref } from "vue";
import { supabase } from "@/auth/supabase";

export function useProfiles() {
    const currentUser = ref(null);

    const fetchCurrentUser = async () => {
        const { data: user, error } = await supabase.auth.getUser();
        if (error) throw error;
        currentUser.value = user.user;
    };

    const getCurrentUserId = async () => {
        if (!currentUser.value) await fetchCurrentUser();
        if (!currentUser.value) throw new Error("No user logged in");

        return currentUser.value.id;
    };

    const checkExistingEmail = async (email) => {
        const res = await api.get(`/profile/check-email/${email}`);
        return res.data.exists;
    };

    return {
        currentUser,
        fetchCurrentUser,
        getCurrentUserId,
        checkExistingEmail,
    };
}
