pub fn add(left: i64, right: i64) -> i64 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;
    use parameterized::parameterized;

    #[parameterized(
        a = {1, 5, 10},
        b = {-1, 5, 99},
        expected = {0, 10, 109}
    )]
    fn test_add(a: i64, b: i64, expected: i64) {
        let result = add(a, b);
        assert_eq!(result, expected);
    }
}
