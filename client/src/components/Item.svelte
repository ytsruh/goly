<script lang="ts">
  import type { Goly } from "../routes/+page";
  export let data: Goly;
</script>

<div class="rounded-lg border border-sky-500 text-sky-500 p-3">
  <p><span class="text-sky-900 text-bold text-lg">Goly:</span> {data.goly}</p>
  <p><span class="text-sky-900 text-bold text-lg">Redirect:</span> {data.redirect}</p>
  <p><span class="text-sky-900 text-bold text-lg">Clicked:</span> {data.clicked}</p>
  <div class="flex justify-between pt-3">
    <button class="bg-sky-500 px-3 py-2 rounded-md text-white">Update</button>
    <button class="bg-sky-900 px-3 py-2 rounded-md text-white">Delete</button>
  </div>
</div>
