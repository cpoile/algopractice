// This is pretty much the instructions in code (clearer than version 2 below):
pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let mut ret: Vec<_> = Vec::new();
    for (r, row) in input.iter().enumerate() {
        for (c, &v) in row.iter().enumerate() {
            if row.iter().all(|&x| v >= x) && input.iter().all(|x| v <= x[c]) {
                ret.push((r, c))
            }
        }
    }
    ret
}

// cleaned up version of 1:
pub fn find_saddle_points_2(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let mut ret: Vec<_> = Vec::new();
    for (r, row) in input.iter().enumerate() {
        for (c, v) in row.iter().enumerate() {
            if v == row.iter().max().unwrap() && v == input.iter().map(|x| &x[c]).min().unwrap() {
                ret.push((r, c))
            }
        }
    }
    ret
}

// old version, lots of waste:
pub fn find_saddle_points_1(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let mut row_max = vec![0u64; input.len()];
    let mut col_min = vec![1u64 << 64 - 1; input[0].len()];

    for (i, r) in input.iter().enumerate() {
        match r.iter().max() {
            Some(m) => row_max[i] = m.clone(),
            None => (),
        }
    }
    for c in 0..input[0].len() {
        for r in input {
            if r[c] < col_min[c] {
                col_min[c] = r[c];
            }
        }
    }

    let mut ret: Vec<_> = Vec::new();
    for (r, row) in input.iter().enumerate() {
        for (c, &v) in row.iter().enumerate() {
            if v == row_max[r] && v == col_min[c] {
                ret.push((r, c))
            }
        }
    }
    ret
}
