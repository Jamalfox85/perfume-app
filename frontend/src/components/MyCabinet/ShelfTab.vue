<template>
    <div>
        <n-button
            @click="addPerfume"
            type="primary"
            style="margin-bottom: 1rem"
        >
            Add Perfume
        </n-button>

        <div class="perfume-grid">
            <div
                v-for="perfume in perfumes"
                :key="perfume.id"
                class="perfume-card"
            >
                <img
                    :src="perfume.image"
                    alt="perfume.name"
                    class="perfume-image"
                />
                <div class="perfume-info">
                    <div class="perfume-name">{{ perfume.name }}</div>
                    <div class="perfume-house">{{ perfume.house }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { defineComponent } from "vue";
import { NButton } from "naive-ui";

export default defineComponent({
    name: "ShelfTab",
    props: ["perfumes"],
    components: {
        NButton,
    },
    methods: {
        addPerfume() {
            // Placeholder for your add perfume functionality
            console.log("Add perfume clicked!");
        },
    },
});
</script>

<style scoped>
.perfume-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
}

.perfume-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
}

.perfume-image {
    width: 100%;
    object-fit: cover;
    border-radius: 0.5rem;
}

.perfume-name {
    font-weight: bold;
    font-size: 1.1rem;
    margin-top: 0.5rem;
}

.perfume-house {
    font-size: 0.9rem;
    color: #666;
}
</style>
