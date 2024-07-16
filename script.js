import tokenizer from 'k6/x/tokenizer';

export default function () {
    let url_or_path = "PATH"; // or "https://path/to/config.json"
    tokenizer.load(url_or_path);

    let prompt = "how are you?";
    let tokens = tokenizer.tokenize(prompt);
    console.log(tokens);
    console.log(tokens.length);
}

