import System.Environment (getArgs)
import System.Exit (exitFailure, exitSuccess)
import Data.Char (isSpace)

main = do
  puzzleInput <- readPuzzleInput
  let result = dialSafe (50, 0, filter (not . null) puzzleInput)
  print result

readPuzzleInput :: IO [String]
readPuzzleInput = do
  args <- getArgs
  case args of
    [filePath] -> do
      contents <- readFile filePath
      return $ lines contents
    [] -> do
      contents <- readFile "../../resources/puzzle-input"
      return $ lines contents
    _  -> do
      putStrLn "Error: too many arguments.\nUsage: program <file>"
      exitFailure

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

trim :: String -> String
trim = f . f
  where f = reverse . dropWhile isSpace