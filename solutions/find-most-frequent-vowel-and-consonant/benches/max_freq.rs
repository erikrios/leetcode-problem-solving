use std::hint::black_box;

use criterion::{Criterion, criterion_group, criterion_main};
use find_most_frequent_vowel_and_consonant::Solution;

fn bench_max_freq(c: &mut Criterion) {
    let s = "abcdefghijklmnopqrstuvwxyz".repeat(100000); // ~2.6M chars

    c.bench_function("one_loop", |b| {
        b.iter(|| Solution::max_freq_sum_one_loop(black_box(s.clone())))
    });

    c.bench_function("two_loops", |b| {
        b.iter(|| Solution::max_freq_sum_two_loops(black_box(s.clone())))
    });
}

criterion_group!(benches, bench_max_freq);
criterion_main!(benches);
