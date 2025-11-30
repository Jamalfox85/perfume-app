<template>
    <div>
        <h3>Your Perfume Collections</h3>
        <p>
            Group your perfumes into collections like “Winter Scents”,
            “Favorites”, etc.
        </p>

        <n-button
            @click="addCollection"
            type="primary"
            style="margin-bottom: 1rem"
        >
            Add Collection
        </n-button>

        <div class="collection-grid">
            <div
                v-for="collection in collections"
                :key="collection.id"
                class="collection-card"
            >
                <div class="collection-name">{{ collection.name }}</div>
                <div class="collection-items">
                    <span
                        v-for="item in collection.perfumes"
                        :key="item.id"
                        class="collection-item"
                        >{{ item.name }}</span
                    >
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { defineComponent } from "vue";
import { NButton } from "naive-ui";

export default defineComponent({
    name: "CollectionsTab",
    props: ["collections"],
    components: {
        NButton,
    },
    methods: {
        addCollection() {
            console.log("Add Collection clicked!");
            this.collections.push({
                id: Date.now(),
                name: "New Collection",
                perfumes: [],
            });
        },
    },
});
</script>

<style scoped>
.collection-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 1rem;
}

.collection-card {
    border: 1px solid #ccc;
    border-radius: 0.5rem;
    padding: 1rem;
}

.collection-name {
    font-weight: bold;
    margin-bottom: 0.5rem;
}

.collection-item {
    background-color: #f0f0f0;
    padding: 0.25rem 0.5rem;
    border-radius: 0.25rem;
    margin-right: 0.25rem;
    display: inline-block;
    margin-bottom: 0.25rem;
}
</style>
