<template>
    <div>
        <h3>Wishlist</h3>
        <p>Perfumes you’re planning to buy.</p>

        <n-button
            @click="addWishlistItem"
            type="primary"
            style="margin-bottom: 1rem"
        >
            Add to Wishlist
        </n-button>

        <div class="wishlist-grid">
            <div v-for="item in wishlist" :key="item.id" class="wishlist-card">
                <img :src="item.image" alt="item.name" class="wishlist-image" />
                <div class="wishlist-info">
                    <div class="wishlist-name">{{ item.name }}</div>
                    <div class="wishlist-house">{{ item.house }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { defineComponent } from "vue";
import { NButton } from "naive-ui";

export default defineComponent({
    name: "WishlistTab",
    props: ["wishlist"],
    components: {
        NButton,
    },
    methods: {
        addWishlistItem() {
            console.log("Add to Wishlist clicked!");
            this.wishlist.push({
                id: Date.now(),
                name: "New Perfume",
                house: "Unknown",
                image: "/images/placeholder.jpg",
            });
        },
    },
});
</script>

<style scoped>
.wishlist-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
}

.wishlist-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
}

.wishlist-image {
    width: 100%;
    object-fit: cover;
    border-radius: 0.5rem;
}

.wishlist-name {
    font-weight: bold;
    font-size: 1.1rem;
    margin-top: 0.5rem;
}

.wishlist-house {
    font-size: 0.9rem;
    color: #666;
}
</style>
