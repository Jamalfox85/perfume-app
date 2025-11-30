<template>
    <div>
        <h3>Create Your Layered Scents</h3>
        <p>
            Add perfumes, body mists, oils, etc., to build your personalized
            scent combinations.
        </p>

        <n-button @click="addLayer" type="primary" style="margin-bottom: 1rem">
            Add Layer
        </n-button>

        <div class="layer-list">
            <div
                v-for="(layeringSet, index) in layeringSets"
                :key="index"
                class="layer-card"
            >
                <div class="layer-name">{{ layeringSet.name }}</div>
                <div class="layer-items">
                    <span
                        v-for="item in layeringSet.items"
                        :key="item.id"
                        class="layer-item"
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
    name: "LayeringTab",
    props: ["layeringSets"],
    components: {
        NButton,
    },
    data() {
        return {
            layers: [
                // Example layer
                {
                    name: "Evening Scent",
                    items: [
                        { id: 1, name: "Chanel No. 5" },
                        { id: 2, name: "Santal 33 Body Oil" },
                    ],
                },
            ],
        };
    },
    methods: {
        addLayer() {
            console.log("Add Layer clicked!");
            // Placeholder for creating a new layer
            this.layers.push({ name: "New Layer", items: [] });
        },
    },
});
</script>

<style scoped>
.layer-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.layer-card {
    border: 1px solid #ccc;
    padding: 1rem;
    border-radius: 0.5rem;
}

.layer-name {
    font-weight: bold;
    margin-bottom: 0.5rem;
}

.layer-item {
    background-color: #f0f0f0;
    padding: 0.2rem 0.5rem;
    border-radius: 0.25rem;
    margin-right: 0.5rem;
    display: inline-block;
    margin-bottom: 0.25rem;
}
</style>
