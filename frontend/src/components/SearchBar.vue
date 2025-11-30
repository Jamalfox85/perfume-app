<template lang="">
    <div class="search-container">
        <n-input
            v-model:value="searchQuery"
            placeholder="Search for perfumes..."
            class="mt-4 w-1/2"
            @input="onSearch"
            @focus="isFocused = true"
            @blur="onBlur"
        />

        <!-- Results -->
        <ul
            v-if="isFocused && (results.length || loading)"
            class="absolute z-10 w-full bg-white border rounded mt-1 max-h-60 overflow-y-auto shadow-lg"
        >
            <li v-if="loading" class="p-2 text-gray-500">Loading...</li>
            <li v-else-if="results.length === 0" class="p-2 text-gray-500">
                No results found
            </li>
            <li
                v-for="perfume in results"
                :key="perfume.id"
                @mousedown.prevent="goToPerfume(perfume.id)"
                class="p-2 hover:bg-gray-100 cursor-pointer"
            >
                {{ perfume.name }} — {{ perfume.house }}
            </li>
        </ul>
    </div>
</template>
<script>
import { NInput } from "naive-ui";
import { usePerfumes } from "@/composables/usePerfumes";
export default {
    components: {
        NInput,
    },
    data() {
        return {
            searchQuery: "",
            results: [],
            loading: false,
            isFocused: false,
            timeout: null,
        };
    },
    methods: {
        async onSearch() {
            clearTimeout(this.timeout);
            this.timeout = setTimeout(async () => {
                if (!this.searchQuery) {
                    this.results = [];
                    return;
                }

                this.loading = true;
                const { searchPerfumes } = usePerfumes();
                try {
                    this.results = await searchPerfumes(this.searchQuery);
                    console.log("Search results:", this.results);
                } catch (err) {
                    console.error("Search error:", err);
                    this.results = [];
                } finally {
                    this.loading = false;
                }
            }, 300); // debounce
        },
        onBlur() {
            setTimeout(() => {
                this.isFocused = false;
            }, 150);
        },
        goToPerfume(id) {
            const router = useRouter();
            router.push({ name: "PerfumeDetails", params: { id } });
            this.query = "";
            this.results = [];
            this.isFocused = false;
        },
    },
};
</script>
<style lang=""></style>
