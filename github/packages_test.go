// Copyright 2021 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "testing"

func TestPackageRegistry_Marshal(t *testing.T) {
	testJSONMarshal(t, &PackageRegistry{}, "{}")

	o := &PackageRegistry{
		AboutURL: String("aurl"),
		Name:     String("name"),
		Type:     String("type"),
		URL:      String("url"),
		Vendor:   String("vendor"),
	}
	want := `{"about_url":"aurl","name":"name","type":"type","url":"url","vendor":"vendor"}`

	testJSONMarshal(t, o, want)
}

func TestPackageFile_Marshal(t *testing.T) {
	testJSONMarshal(t, &PackageFile{}, "{}")

	o := &PackageFile{
		DownloadURL: String("durl"),
		ID:          Int64(1),
		Name:        String("name"),
		SHA256:      String("sha256"),
		SHA1:        String("sha1"),
		MD5:         String("md5"),
		ContentType: String("ct"),
		State:       String("state"),
		Author: &User{
			Login:           String("l"),
			ID:              Int64(1),
			URL:             String("u"),
			AvatarURL:       String("a"),
			GravatarID:      String("g"),
			Name:            String("n"),
			Company:         String("c"),
			Blog:            String("b"),
			Location:        String("l"),
			Email:           String("e"),
			Hireable:        Bool(true),
			Bio:             String("b"),
			TwitterUsername: String("t"),
			PublicRepos:     Int(1),
			Followers:       Int(1),
			Following:       Int(1),
			CreatedAt:       &Timestamp{referenceTime},
			SuspendedAt:     &Timestamp{referenceTime},
		},
		Size:      Int64(1),
		CreatedAt: &Timestamp{referenceTime},
		UpdatedAt: &Timestamp{referenceTime},
	}

	want := `{"download_url":"durl","id":1,"name":"name","sha256":"sha256","sha1":"sha1","md5":"md5","content_type":"ct","state":"state","author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"size":1,"created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `}`

	testJSONMarshal(t, o, want)
}

func TestPackageRelease_Marshal(t *testing.T) {
	testJSONMarshal(t, &PackageRelease{}, "{}")

	o := &PackageRelease{
		URL:             String("url"),
		HTMLURL:         String("hurl"),
		ID:              Int64(1),
		TagName:         String("tn"),
		TargetCommitish: String("tcs"),
		Name:            String("name"),
		Draft:           Bool(true),
		Author: &User{
			Login:           String("l"),
			ID:              Int64(1),
			URL:             String("u"),
			AvatarURL:       String("a"),
			GravatarID:      String("g"),
			Name:            String("n"),
			Company:         String("c"),
			Blog:            String("b"),
			Location:        String("l"),
			Email:           String("e"),
			Hireable:        Bool(true),
			Bio:             String("b"),
			TwitterUsername: String("t"),
			PublicRepos:     Int(1),
			Followers:       Int(1),
			Following:       Int(1),
			CreatedAt:       &Timestamp{referenceTime},
			SuspendedAt:     &Timestamp{referenceTime},
		},
		Prerelease:  Bool(true),
		CreatedAt:   &Timestamp{referenceTime},
		PublishedAt: &Timestamp{referenceTime},
	}

	want := `{"url":"url","html_url":"hurl","id":1,"tag_name":"tn","target_commitish":"tcs","name":"name","draft":true,"author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"prerelease":true,"created_at":` + referenceTimeStr + `,"published_at":` + referenceTimeStr + `}`

	testJSONMarshal(t, o, want)
}

func TestPackageVersion_Marshal(t *testing.T) {
	testJSONMarshal(t, &PackageVersion{}, "{}")

	o := &PackageVersion{
		ID:       Int64(1),
		Version:  String("ver"),
		Summary:  String("sum"),
		Body:     String("body"),
		BodyHTML: String("btnhtml"),
		Release: &PackageRelease{
			URL:             String("url"),
			HTMLURL:         String("hurl"),
			ID:              Int64(1),
			TagName:         String("tn"),
			TargetCommitish: String("tcs"),
			Name:            String("name"),
			Draft:           Bool(true),
			Author: &User{
				Login:           String("l"),
				ID:              Int64(1),
				URL:             String("u"),
				AvatarURL:       String("a"),
				GravatarID:      String("g"),
				Name:            String("n"),
				Company:         String("c"),
				Blog:            String("b"),
				Location:        String("l"),
				Email:           String("e"),
				Hireable:        Bool(true),
				Bio:             String("b"),
				TwitterUsername: String("t"),
				PublicRepos:     Int(1),
				Followers:       Int(1),
				Following:       Int(1),
				CreatedAt:       &Timestamp{referenceTime},
				SuspendedAt:     &Timestamp{referenceTime},
			},
			Prerelease:  Bool(true),
			CreatedAt:   &Timestamp{referenceTime},
			PublishedAt: &Timestamp{referenceTime},
		},
		Manifest:        String("mani"),
		HTMLURL:         String("hurl"),
		TagName:         String("tn"),
		TargetCommitish: String("tcs"),
		TargetOID:       String("tid"),
		Draft:           Bool(true),
		Prerelease:      Bool(true),
		CreatedAt:       &Timestamp{referenceTime},
		UpdatedAt:       &Timestamp{referenceTime},
		PackageFiles: []*PackageFile{
			{
				DownloadURL: String("durl"),
				ID:          Int64(1),
				Name:        String("name"),
				SHA256:      String("sha256"),
				SHA1:        String("sha1"),
				MD5:         String("md5"),
				ContentType: String("ct"),
				State:       String("state"),
				Author: &User{
					Login:           String("l"),
					ID:              Int64(1),
					URL:             String("u"),
					AvatarURL:       String("a"),
					GravatarID:      String("g"),
					Name:            String("n"),
					Company:         String("c"),
					Blog:            String("b"),
					Location:        String("l"),
					Email:           String("e"),
					Hireable:        Bool(true),
					Bio:             String("b"),
					TwitterUsername: String("t"),
					PublicRepos:     Int(1),
					Followers:       Int(1),
					Following:       Int(1),
					CreatedAt:       &Timestamp{referenceTime},
					SuspendedAt:     &Timestamp{referenceTime},
				},
				Size:      Int64(1),
				CreatedAt: &Timestamp{referenceTime},
				UpdatedAt: &Timestamp{referenceTime},
			},
		},
		Author: &User{
			Login:           String("l"),
			ID:              Int64(1),
			URL:             String("u"),
			AvatarURL:       String("a"),
			GravatarID:      String("g"),
			Name:            String("n"),
			Company:         String("c"),
			Blog:            String("b"),
			Location:        String("l"),
			Email:           String("e"),
			Hireable:        Bool(true),
			Bio:             String("b"),
			TwitterUsername: String("t"),
			PublicRepos:     Int(1),
			Followers:       Int(1),
			Following:       Int(1),
			CreatedAt:       &Timestamp{referenceTime},
			SuspendedAt:     &Timestamp{referenceTime},
		},
		InstallationCommand: String("ic"),
	}

	want := `{"id":1,"version":"ver","summary":"sum","body":"body","body_html":"btnhtml","release":{"url":"url","html_url":"hurl","id":1,"tag_name":"tn","target_commitish":"tcs","name":"name","draft":true,"author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"prerelease":true,"created_at":` + referenceTimeStr + `,"published_at":` + referenceTimeStr + `},"manifest":"mani","html_url":"hurl","tag_name":"tn","target_commitish":"tcs","target_oid":"tid","draft":true,"prerelease":true,"created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `,"package_files":[{"download_url":"durl","id":1,"name":"name","sha256":"sha256","sha1":"sha1","md5":"md5","content_type":"ct","state":"state","author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"size":1,"created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `}],"author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"installation_command":"ic"}`

	testJSONMarshal(t, o, want)
}

func TestPackage_Marshal(t *testing.T) {
	testJSONMarshal(t, &Package{}, "{}")

	o := &Package{
		ID:          Int64(1),
		Name:        String("name"),
		PackageType: String("pt"),
		HTMLURL:     String("hurl"),
		CreatedAt:   &Timestamp{referenceTime},
		UpdatedAt:   &Timestamp{referenceTime},
		Visibility:  String("private"),
		Owner: &User{
			Login:           String("l"),
			ID:              Int64(1),
			URL:             String("u"),
			AvatarURL:       String("a"),
			GravatarID:      String("g"),
			Name:            String("n"),
			Company:         String("c"),
			Blog:            String("b"),
			Location:        String("l"),
			Email:           String("e"),
			Hireable:        Bool(true),
			Bio:             String("b"),
			TwitterUsername: String("t"),
			PublicRepos:     Int(1),
			Followers:       Int(1),
			Following:       Int(1),
			CreatedAt:       &Timestamp{referenceTime},
			SuspendedAt:     &Timestamp{referenceTime},
		},
		PackageVersion: &PackageVersion{
			ID:       Int64(1),
			Version:  String("ver"),
			Summary:  String("sum"),
			Body:     String("body"),
			BodyHTML: String("btnhtml"),
			Release: &PackageRelease{
				URL:             String("url"),
				HTMLURL:         String("hurl"),
				ID:              Int64(1),
				TagName:         String("tn"),
				TargetCommitish: String("tcs"),
				Name:            String("name"),
				Draft:           Bool(true),
				Author: &User{
					Login:           String("l"),
					ID:              Int64(1),
					URL:             String("u"),
					AvatarURL:       String("a"),
					GravatarID:      String("g"),
					Name:            String("n"),
					Company:         String("c"),
					Blog:            String("b"),
					Location:        String("l"),
					Email:           String("e"),
					Hireable:        Bool(true),
					Bio:             String("b"),
					TwitterUsername: String("t"),
					PublicRepos:     Int(1),
					Followers:       Int(1),
					Following:       Int(1),
					CreatedAt:       &Timestamp{referenceTime},
					SuspendedAt:     &Timestamp{referenceTime},
				},
				Prerelease:  Bool(true),
				CreatedAt:   &Timestamp{referenceTime},
				PublishedAt: &Timestamp{referenceTime},
			},
			Manifest:        String("mani"),
			HTMLURL:         String("hurl"),
			TagName:         String("tn"),
			TargetCommitish: String("tcs"),
			TargetOID:       String("tid"),
			Draft:           Bool(true),
			Prerelease:      Bool(true),
			CreatedAt:       &Timestamp{referenceTime},
			UpdatedAt:       &Timestamp{referenceTime},
			PackageFiles: []*PackageFile{
				{
					DownloadURL: String("durl"),
					ID:          Int64(1),
					Name:        String("name"),
					SHA256:      String("sha256"),
					SHA1:        String("sha1"),
					MD5:         String("md5"),
					ContentType: String("ct"),
					State:       String("state"),
					Author: &User{
						Login:           String("l"),
						ID:              Int64(1),
						URL:             String("u"),
						AvatarURL:       String("a"),
						GravatarID:      String("g"),
						Name:            String("n"),
						Company:         String("c"),
						Blog:            String("b"),
						Location:        String("l"),
						Email:           String("e"),
						Hireable:        Bool(true),
						Bio:             String("b"),
						TwitterUsername: String("t"),
						PublicRepos:     Int(1),
						Followers:       Int(1),
						Following:       Int(1),
						CreatedAt:       &Timestamp{referenceTime},
						SuspendedAt:     &Timestamp{referenceTime},
					},
					Size:      Int64(1),
					CreatedAt: &Timestamp{referenceTime},
					UpdatedAt: &Timestamp{referenceTime},
				},
			},
			Author: &User{
				Login:           String("l"),
				ID:              Int64(1),
				URL:             String("u"),
				AvatarURL:       String("a"),
				GravatarID:      String("g"),
				Name:            String("n"),
				Company:         String("c"),
				Blog:            String("b"),
				Location:        String("l"),
				Email:           String("e"),
				Hireable:        Bool(true),
				Bio:             String("b"),
				TwitterUsername: String("t"),
				PublicRepos:     Int(1),
				Followers:       Int(1),
				Following:       Int(1),
				CreatedAt:       &Timestamp{referenceTime},
				SuspendedAt:     &Timestamp{referenceTime},
			},
			InstallationCommand: String("ic"),
		},
		Registry: &PackageRegistry{
			AboutURL: String("aurl"),
			Name:     String("name"),
			Type:     String("type"),
			URL:      String("url"),
			Vendor:   String("vendor"),
		},
	}

	want := `{"id":1,"name":"name","package_type":"pt","html_url":"hurl","created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `,"owner":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"package_version":{"id":1,"version":"ver","summary":"sum","body":"body","body_html":"btnhtml","release":{"url":"url","html_url":"hurl","id":1,"tag_name":"tn","target_commitish":"tcs","name":"name","draft":true,"author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"prerelease":true,"created_at":` + referenceTimeStr + `,"published_at":` + referenceTimeStr + `},"manifest":"mani","html_url":"hurl","tag_name":"tn","target_commitish":"tcs","target_oid":"tid","draft":true,"prerelease":true,"created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `,"package_files":[{"download_url":"durl","id":1,"name":"name","sha256":"sha256","sha1":"sha1","md5":"md5","content_type":"ct","state":"state","author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"size":1,"created_at":` + referenceTimeStr + `,"updated_at":` + referenceTimeStr + `}],"author":{"login":"l","id":1,"avatar_url":"a","gravatar_id":"g","name":"n","company":"c","blog":"b","location":"l","email":"e","hireable":true,"bio":"b","twitter_username":"t","public_repos":1,"followers":1,"following":1,"created_at":` + referenceTimeStr + `,"suspended_at":` + referenceTimeStr + `,"url":"u"},"installation_command":"ic"},"registry":{"about_url":"aurl","name":"name","type":"type","url":"url","vendor":"vendor"},"visibility":"private"}`

	testJSONMarshal(t, o, want)
}
