import api from "../services/apiClient";

export function useProfiles() {
    const checkExistingEmail = async (email) => {
        const res = await api.get(`/profile/check-email/${email}`);
        console.log("Email exists:", res.data.exists);
        return res.data.exists;
    };

    return {
        checkExistingEmail,
    };
}
