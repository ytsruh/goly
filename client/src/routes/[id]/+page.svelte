<script lang="ts">
  import { goto } from "$app/navigation";
  import type { Goly } from "./+page";
  import { feedback } from "../../components/FeedbackStore";
  export let data: Goly;
  let showError = false;

  async function handleUpdate(e: any) {
    try {
      data.redirect = e.target[0].value;
      data.goly = e.target[1].value;
      data.random = e.target[1].value === "" ? true : false;

      const response = await fetch("/api/goly", {
        method: "PATCH",
        body: JSON.stringify(data),
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error();
      }
      const res = await response.json();
      feedback.set("Success: Goly has been successfully updated.");
      goto("/");
    } catch (err) {
      feedback.set("Error: An error has occurred. Please try again.");
    }
  }
</script>

<h1 class="text-3xl text-sky-500 my-5 text-center">Goly -- Update</h1>

<div class="flex flex-col mx-auto w-5/6 md:w-1/2 lg:w-1/3 border border-slate-500 rounded-md p-2">
  <form class="min-w-full" on:submit|preventDefault={handleUpdate}>
    <div class="flex flex-col w-full py-2">
      <span>Redirect to</span>
      <input
        type="text"
        class="border border-sky-500 rounded-md p-1 w-full"
        placeholder="https://www.bbc.co.uk"
        value={data.redirect}
        name="redirect"
        required
        autocomplete="off"
      />
    </div>
    <div class="flex flex-col w-full py-2">
      <span>Goly</span>
      <input
        type="text"
        class="border border-sky-500 rounded-md p-1 w-full"
        placeholder="Short link or leave blank to have a random one generated"
        value={data.goly}
        name="goly"
        autocomplete="off"
      />
    </div>
    <div class="py-5">
      <button class="text-white bg-sky-500 rounded-md px-3 py-2 w-full">Update</button>
    </div>
  </form>
</div>
