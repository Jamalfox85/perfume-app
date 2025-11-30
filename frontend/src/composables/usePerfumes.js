import api from "../services/apiClient";

export function usePerfumes() {
    const getPerfume = async (id) => {
        const res = await api.get(`/perfumes/${id}`);
        return res.data;
    };

    const searchPerfumes = async (query) => {
        const res = await api.get("/perfumes", { params: { q: query } });
        return res.data.perfumes;
    };

    return {
        getPerfume,
        searchPerfumes,
    };
}
