<template>
    <n-tabs v-model:value="activeTab" type="line">
        <n-tab-pane name="shelf" tab="My Shelf">
            <ShelfTab :perfumes="perfumes" />
        </n-tab-pane>
        <n-tab-pane name="layering" tab="My Layering Sets">
            <LayeringTab :layeringSets="layeringSets" />
        </n-tab-pane>
        <n-tab-pane name="collections" tab="My Collections">
            <CollectionsTab :collections="collections" />
        </n-tab-pane>
        <n-tab-pane name="wishlist" tab="My Wishlist">
            <WishlistTab :wishlist="wishlist" />
        </n-tab-pane>
    </n-tabs>
</template>

<script>
import { defineComponent } from "vue";
import { NTabs, NTabPane } from "naive-ui";
import { useCabinet } from "@/composables/useCabinet";

import ShelfTab from "@/components/MyCabinet/ShelfTab.vue";
import LayeringTab from "@/components/MyCabinet/LayeringTab.vue";
import CollectionsTab from "@/components/MyCabinet/CollectionsTab.vue";
import WishlistTab from "@/components/MyCabinet/WishlistTab.vue";

export default defineComponent({
    name: "MyCabinet",
    components: {
        NTabs,
        NTabPane,
        ShelfTab,
        LayeringTab,
        CollectionsTab,
        WishlistTab,
    },
    data() {
        return {
            activeTab: "shelf",
            perfumes: [],
            layeringSets: [],
            collections: [],
            wishlist: [],
        };
    },
    mounted() {
        const { fetchCabinet } = useCabinet();
        fetchCabinet().then((data) => {
            this.perfumes = data.shelf;
            this.layeringSets = data.layering_sets;
            this.collections = data.collections;
            this.wishlist = data.wishlist;
            console.log("Cabinet data fetched:", data);
        });
    },
});
</script>

<style scoped></style>
