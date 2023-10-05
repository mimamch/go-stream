const axios = require("axios");

const main = async () => {
  try {
    const res = await axios.get("http://localhost:5000", {
      responseType: "stream",
    });

    res.data.on("data", (chunk) => {
      console.log(chunk.toString());
    });
    res.data.on("end", () => {
      console.log("end");
    });
  } catch (error) {
    console.log(error);
  }
};

main();
