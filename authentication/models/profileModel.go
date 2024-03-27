// package models

// import "errors"

// type Profile struct {
//     UserID   string `json:"userID"`
//     FullName string `json:"fullName"`
//     Email    string `json:"email"`
// }

// var Profiles = []Profile{}

// // Mock database operations
// func AddProfile(p Profile) {
//     Profiles = append(Profiles, p)
// }

// func GetProfile(userID string) (Profile, error) {
//     for _, profile := range Profiles {
//         if profile.UserID == userID {
//             return profile, nil
//         }
//     }
//     return Profile{}, errors.New("profile not found")
// }

// func DeleteProfile(userID string) error {
//     for i, profile := range Profiles {
//         if profile.UserID == userID {
//             Profiles = append(Profiles[:i], Profiles[i+1:]...)
//             return nil
//         }
//     }
//     return errors.New("profile not found")
// }


package models

import "errors"

type Profile struct {
    UserID   string `json:"userID"`
    FullName string `json:"fullName"`
    Email    string `json:"email"`
}

// Initialized with dummy data
var Profiles = []Profile{
    {UserID: "1", FullName: "John Doe", Email: "john.doe@example.com"},
    {UserID: "2", FullName: "Jane Doe", Email: "jane.doe@example.com"},
    {UserID: "3", FullName: "Jim Beam", Email: "jim.beam@example.com"},
}

func AddProfile(p Profile) {
    Profiles = append(Profiles, p)
}

func GetProfile(userID string) (Profile, error) {
    for _, profile := range Profiles {
        if profile.UserID == userID {
            return profile, nil
        }
    }
    return Profile{}, errors.New("profile not found")
}

func DeleteProfile(userID string) error {
    for i, profile := range Profiles {
        if profile.UserID == userID {
            Profiles = append(Profiles[:i], Profiles[i+1:]...)
            return nil
        }
    }
    return errors.New("profile not found")
}

