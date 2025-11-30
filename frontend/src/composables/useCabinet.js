import api from "../services/apiClient";
import { useProfiles } from "./useProfiles";

export function useCabinet() {
    const { getCurrentUserId } = useProfiles();

    const fetchCabinet = async () => {
        const userId = await getCurrentUserId();
        const res = await api.get(`/cabinet/${userId}`);
        return res.data.cabinet;
    };

    return {
        fetchCabinet,
    };
}
