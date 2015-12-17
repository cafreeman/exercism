pub fn anagrams_for<'a>(source: &str, inputs: &[&'a str]) -> Vec<&'a str> {
    // Break the source string into a Vector of chars and sort them for comparison
    let mut source_vec: Vec<char> = source.chars().collect::<Vec<char>>();
    source_vec.sort();
    // sourceVec.sort();
    // Iterate through the inputs array and filter only the elements that satisfy anagram_for
    let res = inputs.iter().filter(|x| anagram_for(&source_vec, x)).cloned().collect();
    // return the Vec of anagrams (or an empty vec)
    res
}

fn anagram_for(sorted_source: &Vec<char>, input: &str) -> bool {
    // If an element is a different length than source word, it can't be an anagram
    if sorted_source.len() != input.len() {
        return false
    } else {
        let mut sorted_input = input.chars().collect::<Vec<char>>();
        sorted_input.sort();
        // create an element-by-element tuple of the two words and then compare them
        let zipped = sorted_source.iter().zip(input.chars());
        if zipped.filter(|&x| *x.0 != x.1).count() > 0 {
            return false;
        }
        return true;
    }
}

fn main() {
    let inputs = ["tan", "stand", "at"];
    let res = anagrams_for("ant", &inputs);
    println!("{}", res.len());
    for i in res {
        println!("{}", i);
    }
}

// fn main() {
//     let input = "tan";
//     let source = "ant";
//     let res = anagram_for(input, source);
// }
