#![feature(inclusive_range_syntax)]

#[derive(PartialEq)]
#[derive(Debug)]
enum Direction {
  North,
  East,
  South,
  West,
}

#[derive(PartialEq)]
struct Point {
  x: i32,
  y: i32
}

fn main() {
  let mut pathstr = String::from("R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, R2, R3, R5, L2, L3, R3, R1, R3");
  //let mut pathstr = String::from("R8, R4, R4, R8");

  pathstr = pathstr.replace(",", "");
  
  let path = pathstr.split_whitespace();
  
  let mut x:i32 = 0;
  let mut y:i32 = 0;
  let mut dir = Direction::North;
  let mut visited: Vec<Point> = Vec::new();
  let mut hq: Point = Point{x: 0, y: 0};
  let mut found = false;
  
  for segment in path {
    //println!("{}", segment);
    dir = turn(dir, segment.chars().nth(0).unwrap());
    let blocks:i32 = segment.to_string()[1..].parse::<i32>().unwrap();
    for _ in 0..blocks
    {
        match dir {
          Direction::North => {
            y += 1;
          }
          Direction::East => {
            x += 1;
          }
          Direction::South => {
            y -= 1;
          }
          Direction::West => {
            x -= 1;
          }
        }
        //println!("{} {}", x, y);
        if !found {
            for point in &visited {
              if x == point.x && y == point.y {
                //println!("Found! {} {}", x, y);
                found = true;
                hq = Point{x: x, y: y}; 
              }
            }
            
            visited.push(Point{x: x, y: y});
        }
    }
    
    //println!("X: {}, Y: {}", x, y);
    //println!("visited: {}", visited.len());
  }
  
  println!("location: x: {} y: {}", x, y);
  println!("distance: {}", x.abs() + y.abs());
  println!("HQ location: x: {} y: {}", hq.x, hq.y);
  println!("HQ distance: {}", hq.x.abs() + hq.y.abs());
}

fn turn(facing: Direction, dir: char) -> Direction {
  let mut newfacing = Direction::North;
  
  match facing {
    Direction::North => {
      if dir == 'R' {
        newfacing = Direction::East;
      } else if dir == 'L' {
        newfacing = Direction::West;
      }
    }
    Direction::East => {
      if dir == 'R' {
        newfacing = Direction::South;
      } else if dir == 'L' {
        newfacing = Direction::North;
      }
    }
    Direction::South => {
      if dir == 'R' {
        newfacing = Direction::West;
      } else if dir == 'L' {
        newfacing = Direction::East;
      }
    }
    Direction::West => {
      if dir == 'R' {
        newfacing = Direction::North;
      } else if dir == 'L' {
        newfacing = Direction::South;
      }
    }
  }
  
  newfacing
}