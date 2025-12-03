import AocUtils

fallbackPath :: String
fallbackPath = "/home/sca/Programming/advent-of-code/2025/Day-1/resources/puzzle-input"

solutionFormat :: String
solutionFormat = ">>> The solution is: "

main :: IO ()
main = do
  puzzleInput <- readPuzzleInput fallbackPath lines
  let (_, result, _) = dialSafe (50, 0, filter (not . null) puzzleInput)
  putStrLn (solutionFormat ++ show result)

-- (current dial position, times dial hit 0, instruced actions for the dial)
dialSafe :: (Integer, Integer, [String]) -> (Integer, Integer, [String])
dialSafe (dialPos, timesDialHitZero, []) = (dialPos, timesDialHitZero, [])
dialSafe (dialPos, timesDialHitZero, action:rest) = do
  let dialPos' = dialOnce dialPos action
  let timesDialHitZero' = if dialPos' == 0 then timesDialHitZero + 1 else timesDialHitZero
  dialSafe (dialPos', timesDialHitZero', rest)

dialOnce :: Integer -> String -> Integer
dialOnce dialPos action = do
  let dir = getDirInt $ head action
  let steps = read (trim (drop 1 action))
  let dial' = dialPos + dir * steps
  abs (if (dial' < 0) || (dial' > 99) then dial' `mod` 100 else dial')

getDirInt :: Char -> Integer
getDirInt 'L' = -1
getDirInt 'R' = 1
getDirInt _ = 0
