package main

func judgeCircle(moves string) bool {
    if len(moves) == 0 {
        return true
    }

    if len(moves) == 1 {
        return false
    }

    lr, ud := 0, 0
    for i := 0; i < len(moves); i ++ {
        switch string(moves[i]) {
            case "R":
                lr ++
            case "L":
                lr --
            case "U":
                ud ++
            case "D":
                ud --
            default:
        }
    }
    if lr != 0 || ud != 0 {
        return false
    }
    return true
}