<script lang="ts">
  import type { Goly } from "../routes/+page";
  import Card from "./Card.svelte";
  export let data: Goly;
</script>

<Card>
  <p><span class="text-sky-900 text-bold text-lg">Goly:</span> {data.goly}</p>
  <p><span class="text-sky-900 text-bold text-lg">Redirect:</span> {data.redirect}</p>
  <p><span class="text-sky-900 text-bold text-lg">Clicked:</span> {data.clicked}</p>
  <div class="flex justify-between pt-3">
    <button class="bg-white px-3 py-2 rounded-md text-sky-900">Update</button>
    <button class="bg-sky-900 px-3 py-2 rounded-md text-white">Delete</button>
  </div>
</Card>
