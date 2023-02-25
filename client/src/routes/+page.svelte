<script lang="ts">
  import type { Goly } from "./+page";
  type PageData = {
    golies: Goly[];
  };
  export let data: PageData;
  import Item from "../components/Item.svelte";
</script>

<h1 class="text-3xl text-sky-500 my-5 text-center">Goly -- The Go URL Shortener</h1>
<div class="py-2 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5 px-10 align-center">
  {#each data.golies as goly}
    <Item data={goly} />
  {/each}
</div>
