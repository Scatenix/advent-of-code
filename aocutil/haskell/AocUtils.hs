module AocUtils
(
    readPuzzleInput,
    trim
) 
where

import System.Environment (getArgs)
import System.Exit (exitFailure)
import Data.Char (isSpace)

-- FallbackPath, function (contents -> output) -> IO output
readPuzzleInput :: String -> (String -> a) -> IO a
readPuzzleInput fallbackPath parser = do
  args <- getArgs
  case args of
    [filePath] -> do
      contents <- readFile filePath
      return $ parser contents
    [] -> do
      contents <- readFile fallbackPath
      return $ parser contents
    _  -> do
      putStrLn "Error: too many arguments.\nUsage: program <file>"
      exitFailure

trim :: String -> String
trim = f . f
  where f = reverse . dropWhile isSpace